package handlers

import (
	"github.com/gin-gonic/gin"
)

// SampleHandler provides a simple health check endpoint returning "Hello World!"
func SampleHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
