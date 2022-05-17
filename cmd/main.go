package main

import (
	database "go-workspace/src/mssgs-crud/db/database"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'v got it",
	})
}

func PostPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'v posted it",
	})
}

func UpdatePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'v posted it",
	})
}

func DelPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'v erased it",
	})
}

func main() {
	r := gin.Default()

	database.Init()

	r.GET("/", HomePage)
	r.POST("/", PostPage)
	r.PATCH("/", UpdatePage)
	r.DELETE("/", DelPage)
	r.Run()
}
