package service

import (
	"log"
	constant "omnial-simulator/constant"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDbInstance() (*gorm.DB, error) {

	dbms := os.Getenv("DBMS")

	// Create a DB DSN
	dsn := constant.MakeDsn()

	// connect to database
	db, err := gorm.Open(dbms, dsn)

	if err != nil {
		log.Println("DB connection failed")
		log.Println(err)
		return nil, err
	}

	return db, nil
}
