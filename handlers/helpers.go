package handlers

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	constant "omnial-simulator/constant"
)

// Generates a random string of specified length
func GenerateRandomString(length int, charType string) (string, error) {
	var charset string
	if charType == "alpha" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if charType == "alphaUpper" {
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if charType == "alphaLower" {
		charset = "abcdefghijklmnopqrstuvwxyz"
	} else if charType == "num" {
		charset = "0123456789"
	} else if charType == "alphaNum" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	} else if charType == "alphaUpperNum" {
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	} else if charType == "alphaLowerNum" {
		charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	} else if charType == "alphaNumSym" {
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;:'\",.<>?`~"
	}

	charsetLength := big.NewInt(int64(len(charset)))

	var builder strings.Builder
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		randomChar := charset[randomIndex.Int64()]
		builder.WriteByte(randomChar)
	}

	return builder.String(), nil
}

func processMap(responseData interface{}, req interface{}) {

	switch data := responseData.(type) {
	case map[string]interface{}:
		// Process a single map
		replaceStringProcess(data, req)
	case []interface{}:
		// Process each map in the array
		for _, item := range data {
			if mapItem, ok := item.(map[string]interface{}); ok {
				replaceStringProcess(mapItem, req)
			} else {
				log.Println("Expected array item to be of type map[string]interface{}, got another type")
			}
		}
	default:
		log.Println("Unexpected type for responseData")
	}
}

// Replacement item processing
func replaceStringProcess(responseData map[string]interface{}, req interface{}) {

	for key, value := range responseData {
		if req != nil {
			reqMap := req.(map[string]interface{})
			if value == constant.Replace {
				// Processing items that are requests
				if reqValue, exists := reqMap[key]; exists {
					responseData[key] = reqValue
					continue
				}
				log.Printf("The value for %s does not exist in the request\n", key)

			}
		}
		if value == constant.Nowtime {
			//　current time
			responseData[key] = time.Now().Format(time.RFC3339)
			continue
		}

		if strValue, ok := value.(string); ok {

			re := regexp.MustCompile(`rand\((\d+), (\w+)`)
			if re.MatchString(strValue) {
				match := re.FindStringSubmatch(strValue)

				numberStr := match[1]
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					log.Printf("fail to transform into number")
					continue
				}

				charType := strings.Join(strings.Fields(match[2]), "")

				randomValue, _ := GenerateRandomString(number, charType)
				responseData[key] = randomValue
				continue

			}
		}
	}
}

// Common processing for request failure
func FailedRequest(w http.ResponseWriter, faliMessage string, httpstatus int) {
	data := map[string]interface{}{
		"code":    httpstatus,
		"message": faliMessage,
	}
	failJson, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(failJson)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(failJson)
}

func extractEndpointFromURL(method string, path string) string {

	// Find the position of the last slash using strings.LastIndex
	lastSlashIdx := strings.LastIndex(path, "/")

	if lastSlashIdx >= 0 {

		substr := path[lastSlashIdx+1:]

		// Use strings.Contains to determine if it contains a colon
		if colonIdx := strings.Index(substr, ":"); colonIdx >= 0 {
			// If a colon is included, return everything after the colon
			return substr[colonIdx+1:]
		}

		// If the colon is not included, everything after the slash is returned unchanged
		return substr
	}

	// If the slash is not found TODO
	return path

}
