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
	router.GET("/users/:id", handlers.GetUserByIdOrName(db))
	// router.GET("/users/name", handlers.GetUserByName(db))
	router.POST("/users", handlers.AddUser(db))
	router.DELETE("/users/:id", handlers.DeleteUser(db))

	router.GET("/messages", handlers.GetMessages(db))
	router.POST("/messages", handlers.AddMessage(db))
	router.DELETE("/messages/:id", handlers.DeleteMessage(db))

	router.Run()
}
