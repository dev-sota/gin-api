package controllers

import "github.com/gin-gonic/gin"

// Index action: GET /
func Index(c *gin.Context) {
	result := gin.H{
		"message": "GET",
	}
	c.JSON(200, result)
}
