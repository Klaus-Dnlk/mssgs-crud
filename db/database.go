package db

import (
	"mssgs-crud/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=crud port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database!")
	}

	db.AutoMigrate(&models.User{}, &models.Message{})
	return db
}
