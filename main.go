package main

import (
	"mssg-crud/Handlers"
	DB "mssgs-crud/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := DB.Connection()

	router := gin.Default()

	router.GET("/home", Handlers.GetUsers(db))
	router.POST("/post", Handlers.AddUser(db))

	router.Run()
}
