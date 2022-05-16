package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

db, err := gorm.Open(postgres.New(postgres.Config{
	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})



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
	r.GET("/", HomePage)
	r.POST("/", PostPage)
	r.PATCH("/", UpdatePage)
	r.DELETE("/", DelPage)
	r.Run()
}
