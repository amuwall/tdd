package handler

import (
	"github.com/gin-gonic/gin"
	"tdd/web/middleware"
)

func Register(engine *gin.Engine) {
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recovery())

	apiV1 := engine.Group("/api/v1")
	{
		apiV1.GET("/hello", Hello)
		apiV1.GET("/users", GetUsers)
		apiV1.POST("/users/search", SearchUsers)
	}
}
