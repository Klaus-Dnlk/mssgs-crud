package handlers

import (
	"mssgs-crud/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var users Models.User

		result, err := db.Find(&users) // cannot initialize 2 variables with 1 values
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"qwerqwer: ": result,
		})
	}

	// return func(c *gin.Context) {
	// 	var users []Models.User

	// 	// Get all records
	// 	result, err := DB.Find(&users)

	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, result)
	// }

}

func AddUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	}
}
