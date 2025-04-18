package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func ConnectDatabase() {
	var err error
	dsn := "user=gorm dbname=gorm password=gorm host=localhost port=5444 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database")
	}
}
