package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	service "omnial-simulator/service"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) (bool, string, int) {

	// Get the path (API name)

	endpoint := r.URL.Path
	log.Println("endpoint: ", endpoint)

	var req interface{}
	var responseJson string
	var errMsg string
	var resErr error
	var httpStatus int
	var sleepTime int

	foundResponse := false

	// Decode the request body
	ErrDec := json.NewDecoder(r.Body).Decode(&req)

	if ErrDec != nil && ErrDec != io.EOF {
		log.Println(ErrDec)
		errMsg = ErrDec.Error()
		return false, errMsg, 404
	}

	jsonReq, _ := json.Marshal(req)

	// Print RLOG
	log.Printf("[RLOG] Receive a request Method: %s, URL: %s, Header: %v, Body: %s\n",
		r.Method, r.URL.String(), r.Header, jsonReq)

	// path-> endpoint
	leaf := extractEndpointFromURL(r.Method, endpoint)

	// get ApiInfoKey
	apiInfoList, err := service.GetAPIKeyInfoListByTargetAPI(leaf)

	if err != nil {
		errMsg = "Failed to get ApiInfoKey"
		return false, errMsg, 404
	}

	if len(apiInfoList) == 0 {
		errMsg = "endpoint「" + endpoint + "」does not exist in the DB"
		return false, errMsg, 404
	}

loopApiInfo:
	for _, apiInfo := range apiInfoList {

		keyInfoList, errGetItem := service.GetKeyInfoKeyAndItemListByApiInfoKey(apiInfo.ApiInfoKey)

		if errGetItem != nil {
			errMsg = "Failed to get keyInfo"
			return false, errMsg, 404
		}

	loopKeyInfo:
		for _, keyInfo := range keyInfoList {
			keyItem := keyInfo.KeyItem
			itemSection := keyInfo.ItemSection
			keyInfoKey := keyInfo.KeyInfoKey

			var keyItemValue string
			var okString bool
			// Different handling based on ItemsSection
			if itemSection == "HEADER" {
				// Get headers from request and build response
				headers := r.Header
				keyHeaderValue := headers.Get(keyItem)
				if keyHeaderValue != "" {
					log.Println("keyHeaderValue:", keyHeaderValue)
					// Get the responseJson stored in the database
					responseJson, httpStatus, sleepTime, resErr = service.GetReTelegram(keyInfoKey, keyHeaderValue)
					if resErr == nil {
						foundResponse = true
						break loopApiInfo
					}

				} else {
					errMsg = "Item does not exist in header"
					return false, errMsg, 404
				}

			} else if itemSection == "BODY" {
				if ErrDec == io.EOF {
					log.Println("BODY is none")
					continue loopKeyInfo
				}
				//  Determining the existence of "." in KeyItem (ex: "amount.value")
				dotIndex := strings.Index(keyItem, ".")

				if dotIndex == -1 {

					//　"." if not present

					keyItemValue, okString = req.(map[string]interface{})[keyItem].(string)
					keyItemValueArray, okArray := req.(map[string]interface{})[keyItem].([]interface{})

					if okArray {
						var stringValues []string
						for _, elem := range keyItemValueArray {
							// Convert each element to a string using type assertions.
							if stringValue, ok := elem.(string); ok {
								stringValues = append(stringValues, stringValue)
							}
						}

						// Join an array of strings into a single string using strings.Join()
						keyItemValue = strings.Join(stringValues, ", ")

					} else if !okString {
						log.Println("keyItemValue acquisition failed")
						continue loopKeyInfo
					}

				} else {

					// Split the key string into multiple parts
					keyParts := strings.Split(keyItem, ".")

					// Search from outer JSON
					currentItem := req.(map[string]interface{})

					// iterate through all parts except the last part
					for _, part := range keyParts[:len(keyParts)-1] {
						if nextItem, ok := currentItem[part].(map[string]interface{}); ok {
							currentItem = nextItem
						} else {
							// If the intermediate key does not exist or the types do not match, log an error and continue
							log.Printf("Key %s not found or is not a JSON object\n", part)
							continue loopKeyInfo
						}
					}

					// Get the value corresponding to the last key
					finalValue := currentItem[keyParts[len(keyParts)-1]]

					// Processing according to value type
					if numValue, okNum := finalValue.(float64); okNum {
						keyItemValue = fmt.Sprintf("%.0f", numValue)
					} else if stringValue, okString := finalValue.(string); okString {
						keyItemValue = stringValue
					} else {
						log.Println("get keyItemValue error: value is not a number or a string")
						continue loopKeyInfo
					}

				}

				// Get the responseJson stored in the database
				responseJson, httpStatus, sleepTime, resErr = service.GetReTelegram(keyInfoKey, keyItemValue)
				if resErr == nil {
					foundResponse = true
					log.Println("KeyItem: ", keyItem)
					log.Println("ItemSection: ", itemSection)
					log.Println("KeyInfoKey: ", keyInfoKey)
					log.Println("keyItemValue: ", keyItemValue)
					break loopApiInfo
				}

			} else if itemSection == "PATHPARAMATER" {
				parts := strings.Split(endpoint, "/")
				if len(parts) == 0 {
					errMsg = "url error format"
					return false, errMsg, 500
				}
				lastSegment := parts[len(parts)-1]

				if leaf == "getTransaction" {
					lastSlashIdx := strings.LastIndex(endpoint, "/")
					PATHPARAMATER := endpoint[lastSlashIdx+1:]
					log.Print("pathparamater:", PATHPARAMATER)

					responseJson, httpStatus, sleepTime, resErr = service.GetReTelegram(keyInfoKey, PATHPARAMATER)
					if resErr == nil {
						foundResponse = true
						break loopApiInfo
					}
				}
				if strings.HasSuffix(lastSegment, leaf) {
					PATHPARAMATER := strings.TrimSuffix(lastSegment, ":"+leaf)

					log.Print("PATHPARAMATER:", PATHPARAMATER)

					responseJson, httpStatus, sleepTime, resErr = service.GetReTelegram(keyInfoKey, PATHPARAMATER)
					if resErr == nil {
						foundResponse = true
						break loopApiInfo
					}
				}

			} else if itemSection == "QUERYPARAMETER" {
				queryParam := r.URL.Query().Get(keyItem)
				responseJson, httpStatus, sleepTime, resErr = service.GetReTelegram(keyInfoKey, queryParam)
				if resErr == nil {
					foundResponse = true
					break loopApiInfo
				}

			} else if itemSection == "" {
				responseJson, httpStatus, sleepTime, resErr = service.GetReTelegramBykeyInfoKey(keyInfoKey)
				if resErr == nil {
					foundResponse = true
					break loopApiInfo
				}
			} else {
				errMsg = "ItemsSection does not exist"
				return false, errMsg, 404
			}
		}

	}

	// If no matching response is found
	if !foundResponse {
		errMsg = "ReTelegramの取得に失敗しました"
		return false, errMsg, 404
	}
	// Sleep
	time.Sleep(time.Duration(sleepTime) * time.Second)

	// Set the value for transactionId in responseJson
	var responseData interface{}

	// Array judgment
	if strings.HasPrefix(string(responseJson), "[") {
		var arrData []interface{}
		if err := json.Unmarshal([]byte(responseJson), &arrData); err != nil {
			log.Printf("Failed to unmarshal JSON array: %v", err)
			return false, err.Error(), 500
		}
		responseData = arrData
	} else {
		var responseDataMap map[string]interface{}
		if err := json.Unmarshal([]byte(responseJson), &responseDataMap); err != nil {
			log.Printf("Failed to unmarshal JSON object: %v", err)
			return false, err.Error(), 500
		}
		responseData = responseDataMap
	}

	// Set the value from request
	processMap(responseData, req)

	// Convert responseData to JSON string
	modifiedResponseJson, _ := json.Marshal(responseData)

	// Check if ModifiedResponseJson represents a JSON array
	if strings.HasPrefix(string(modifiedResponseJson), "[") {
		var arrData []interface{}
		if err := json.Unmarshal([]byte(modifiedResponseJson), &arrData); err != nil {
			log.Fatalf("Failed to unmarshal JSON array: %v", err)
			return false, err.Error(), 500
		}

	} else {
		var responseDataMap map[string]interface{}
		if err := json.Unmarshal([]byte(modifiedResponseJson), &responseDataMap); err != nil {
			log.Fatalf("Failed to unmarshal JSON object: %v", err)
			return false, err.Error(), 500
		}

	}

	// Convert modifiedResponseJson to json.RawMessage
	rawMessage := json.RawMessage(modifiedResponseJson)

	// returns rawMessage
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(rawMessage)

	// Print SLOG
	log.Printf("[SLOG] Send a response Method: %v, URL: %v, StatusCode: %v, Header: %v, Body: %v\n",
		r.Method, r.URL.String(), httpStatus, w.Header(), string(modifiedResponseJson))

	return true, "", 200

}
