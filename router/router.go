package router

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter
func SetupRouter() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.GET("/", controllers.ListUsers)
	}

	return router
}
