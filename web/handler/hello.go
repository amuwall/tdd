package handler

import (
	"github.com/gin-gonic/gin"
	"tdd/web/response"
)

func Hello(c *gin.Context) {
	response.Success(c, gin.H{
		"hello": "world",
	})
}
