package controller

import (
	"github.com/gin-gonic/gin"
)

/**
handler 404 not found Router
 */
func NotFound(context *gin.Context) {
	context.JSON(404, gin.H{
		"error": "404 not found",
		"url":   context.Request.URL,
	})
}
