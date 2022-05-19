package main

import (
	"mssgs-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	DB.Connection()

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/users", handlers.GetUsers)
	}

	r.Run()

	// router := gin.Default()

	// router.GET("/home", Handlers.GetUsers(db))
	// router.POST("/post", Handlers.AddUser(db))

	// router.Run()
}
