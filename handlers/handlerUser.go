package handlers

import (
	"mssgs-crud/Models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []Models.User

		err := db.Find(&users).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user []Models.User
		id := c.Params.ByName("id")
		name := c.Params.ByName("name")
		val, err := strconv.Atoi(id)

		if err != nil{
			err := db.Preload("SentMessages").Preload("ReceivedMessages").Find(&user, Models.User{Name: name}).Error
			if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, user)
		} else {
			err := db.Preload("SentMessages").Preload("ReceivedMessages").First(&user, val).Error
			if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
		}
	}
}

// func GetUserByName(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user []Models.User
// 		name := c.Params.ByName("name")

// 		err := db.Preload("SentMessages").Preload("ReceivedMessages").Find(&user, Models.User{Name: name}).Error
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 		c.JSON(http.StatusOK, user)
// 	}
// }

func AddUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser Models.User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := db.Create(&newUser).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"User created": newUser,
		})
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		
		var users []Models.User

		err := db.Delete(&users, id).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, "User deleted")
	}
}
