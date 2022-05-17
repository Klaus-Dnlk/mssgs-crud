package database

import (
	"fmt"
	"log"
	"mssgs-crud/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=crud port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Users{}, &models.Msgs{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("database is unavailable. Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
