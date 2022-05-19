package db

import (
	"mssgs-crud/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connection() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=crud port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database!")
	}

	db.AutoMigrate(&Models.User{}, &Models.Message{})
	// return db
}

func GetDB() *gorm.DB {
	return db
}
