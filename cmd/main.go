package main

import (
	database "mssgs-crud/db"
	"mssgs-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var db = database.Init()
	r := gin.Default()

	r.GET("/home", handlers.HomePage)
	r.POST("/send", handlers.PostPage)
	// r.PATCH("/", UpdatePage)
	// r.DELETE("/del", DelPage)
	r.Run()
}
