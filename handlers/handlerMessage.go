package handlers

import (
	"mssgs-crud/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMessages(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var users []Models.User

		err := db.Find(&users).Error // cannot initialize 2 variables with 1 values
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, users)
	}
}

func AddMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := db.Create(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	}
}
