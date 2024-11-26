package http

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"Error": message,
	})
}
