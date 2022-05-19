package main

import (
	"mssgs-crud/db"
	"mssgs-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connection()

	router := gin.Default()

	router.GET("/home", handlers.GetUsers)
	router.POST("/post", handlers.AddUser)

	router.Run()
}
