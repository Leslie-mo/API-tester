package service

import (
	models "omnial-simulator/models/dto"
)

// select
func GetReTelegram(keyInfoKey string, keyValue string) (string, int, int, error) {
	// Create a DB instance
	db, err := GetDbInstance()
	if err != nil {
		return "DB connection failed:", 0, 0, err
	}

	// Create a Response object
	response := &models.Response{}

	// Get response based on API name and key value
	err = db.Where(" convert(KEY_INFO_KEY USING utf8) = ? AND convert(KEY_VALUE USING utf8mb4) collate utf8mb4_bin = ?", keyInfoKey, keyValue).Order("UPDATE_TIME DESC").First(response).Error

	if err != nil {

		return "Response message acquisition failure", 0, 0, err

	}

	return response.ReTelegram, response.HttpStatus, response.SleepTime, nil
}

func GetReTelegramBykeyInfoKey(keyInfoKey string) (string, int, int, error) {
	// Create a DB instance
	db, err := GetDbInstance()
	if err != nil {
		return "DB connection failed.", 0, 0, err
	}

	// Create a Response object
	response := &models.Response{}

	// Get response based on API name and key value
	err = db.Where(" convert(KEY_INFO_KEY USING utf8) = ? ", keyInfoKey).Order("UPDATE_TIME DESC").First(response).Error

	if err != nil {

		return "Response message acquisition failure", 0, 0, err

	}

	return response.ReTelegram, response.HttpStatus, response.SleepTime, nil
}
