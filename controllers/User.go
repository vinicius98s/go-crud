package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUsers
func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Vinicius",
	})
}
