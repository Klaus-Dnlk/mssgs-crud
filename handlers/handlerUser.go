package handlers

import (
	"fmt"
	"io/ioutil"
	"mssgs-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	// Get all records
	result := db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		result,
	})
}

func AddUser(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}
