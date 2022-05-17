package database

import (
	"fmt"
	"log"
	"mssgs-crud/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=crud port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Users{}, &models.Messages{})
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		db = Init()
		var sleep = time.Duration(1)
		for db == nil {
			sleep = sleep * 2
			fmt.Printf("database is unavailable. Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			db = Init()
		}
	}
	return db
}
