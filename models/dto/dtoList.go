package models

import "time"

// APIInfo
type APIInfo struct {
	ApiInfoKey string    `gorm:"column:API_INFO_KEY;type:varchar(50);primaryKey"`
	TargetAPI  string    `gorm:"column:TARGET_API;type:varchar(50)"`
	CreateTime time.Time `gorm:"column:CREATE_TIME;type:timestamp;default:current_timestamp"`
	UpdateTime time.Time `gorm:"column:UPDATE_TIME;type:timestamp;default:current_timestamp on update current_timestamp"`
}

// KeyInfo
type KeyInfo struct {
	KeyInfoKey  string    `gorm:"column:KEY_INFO_KEY;type:varchar(50);primaryKey"`
	ApiInfoKey  string    `gorm:"column:API_INFO_KEY;type:varchar(50)"`
	KeyItem     string    `gorm:"column:KEY_ITEM;type:varchar(50)"`
	ItemSection string    `gorm:"column:ITEM_SECTION;type:varchar(50)"`
	CreateTime  time.Time `gorm:"column:CREATE_TIME;type:timestamp;default:current_timestamp"`
	UpdateTime  time.Time `gorm:"column:UPDATE_TIME;type:timestamp;default:current_timestamp on update current_timestamp"`
}

// Response
type Response struct {
	ResponseKey string    `gorm:"column:RESPONSE_KEY;type:varchar(50);primaryKey"`
	KeyInfoKey  string    `gorm:"column:KEY_INFO_KEY;type:varchar(50);primaryKey"`
	KeyValue    string    `gorm:"column:KEY_VALUE;type:varchar(100)"`
	HttpStatus  int       `gorm:"column:HTTPSTATUS"`
	ReTelegram  string    `gorm:"column:RE_TELEGRAM;type:varchar(2000)"`
	SleepTime   int       `gorm:"column:SLEEP_TIME"`
	CreateTime  time.Time `gorm:"column:CREATE_TIME;type:timestamp;default:current_timestamp"`
	UpdateTime  time.Time `gorm:"column:UPDATE_TIME;type:timestamp;default:current_timestamp on update current_timestamp"`
}

func (APIInfo) TableName() string {
	return "API_INFO"
}

func (KeyInfo) TableName() string {
	return "KEY_INFO"
}

func (Response) TableName() string {
	return "RESPONSE"
}
