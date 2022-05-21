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
	router.GET("/users/:id", handlers.GetUserById(db))
	router.POST("/users", handlers.AddUser(db))
	router.DELETE("/users/:id", handlers.DeleteUser(db))

	router.GET("/messages", handlers.GetMessages(db))
	router.POST("/messages", handlers.AddMessage(db))
	router.DELETE("/message/:id", handlers.DeleteMessage(db))

	router.Run()
}
