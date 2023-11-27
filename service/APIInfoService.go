package service

import (
	"fmt"
	"log"
	models "omnial-simulator/models/dto"
)

func GetAPIKeyInfoListByTargetAPI(targetAPI string) ([]models.APIInfo, error) {
	// Create a DB instance
	db, err := GetDbInstance()
	if err != nil {
		return nil, fmt.Errorf("DB connection failed: %v", err)
	}

	// Create an APIInfoList object
	var apiInfoList []models.APIInfo
	err = db.Where("TARGET_API LIKE ?", "%"+targetAPI).Order("UPDATE_TIME DESC").Find(&apiInfoList).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return apiInfoList, nil
}
