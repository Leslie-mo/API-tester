package service

import (
	"log"
	models "omnial-simulator/models/dto"
)

func GetKeyInfoKeyAndItemListByApiInfoKey(apiInfoKey string) ([]models.KeyInfo, error) {
	// Create a DB instance
	db, err := GetDbInstance()
	if err != nil {
		return nil, err
	}

	// Create a KeyInfo object
	var keyInfoList []models.KeyInfo

	// Select("KEY_INFO_KEY")
	err = db.Where("API_INFO_KEY = ?", apiInfoKey).Order("UPDATE_TIME DESC").Find(&keyInfoList).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return keyInfoList, nil
}
