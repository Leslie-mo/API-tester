package constant

import (
	"os"
)

const (
	// replace string
	Replace = "{replace}"
	// replacemethod: nowtime
	Nowtime = "nowTime()"
)

type DatabaseConfig struct {
	Username string `ini:"Username"`
	Password string `ini:"Password"`
	DBName   string `ini:"DBName"`
	DBPort   string `ini:"DBPort"`
}

var dbConfig DatabaseConfig

type PortConfig struct {
	PortValue string `ini:"Port"`
}

var Port string

// make up DSN
func MakeDsn() string {
	dbConfig = DatabaseConfig{
		Username: os.Getenv("USER_NAME"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		DBPort:   os.Getenv("DB_PORT"),
	}

	dsn := dbConfig.Username + ":" + dbConfig.Password + "@" + dbConfig.DBPort + "/" + dbConfig.DBName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	return dsn
}
