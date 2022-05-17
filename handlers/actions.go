package handlers

import (
	"mssgs-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method":  http.MethodGet,
			"code":    http.StatusOK,
			"message": "I'v got it",
		},
	})
}

func PostPage(c *gin.Context) {

	var user *models.Users
	// var message *models.Messages

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"bad_request"})
		return
	}

	db = append(db, user_name)

}

// func UpdatePage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "I'v posted it",
// 	})
// }

// func DelPage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "I'v erased it",
// 	})
// }
