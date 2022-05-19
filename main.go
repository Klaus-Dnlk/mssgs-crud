package main

import (
	DB "mssgs-crud/db"
	"mssgs-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db := DB.Connection()

	router := gin.Default()

	router.GET("/users", handlers.GetUsers(db))
	router.POST("/users", handlers.AddUser(db))

	router.Run()
}
