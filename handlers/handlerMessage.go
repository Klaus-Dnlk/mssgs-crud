package handlers

import (
	"mssgs-crud/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMessages(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var messages []Models.Message

		err := db.Find(&messages).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, messages)
	}
}

func AddMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message Models.Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := db.Create(&message).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"Message content": message,
		})
	}
}

func DeleteMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")

		var message []Models.Message

		err := db.Delete(&message, id).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, "Message deleted")
	}
}
