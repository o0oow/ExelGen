package db

import (
	"github.com/cengsin/oracle"
	"gorm.io/gorm"
	"log"
	"sportsman/configs"
)

var db *gorm.DB

func CloseDB() {
	sqlDb, err := db.DB()
	sqlDb.Close()
	if err != nil {
		log.Println("Error", err)
	}
}
func Setup() {
	var err error
	db, err = gorm.Open(oracle.Open(configs.Config.DB))
	if err != nil {
		log.Println("db.Setup err", err)
	}
}
func GetDB() *gorm.DB {

	return db
}
