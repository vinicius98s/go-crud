package router

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter makes all routes available
func SetupRouter() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.GET("/", controllers.ListUsers)
		users.POST("/", controllers.CreateUser)
	}

	return router
}
