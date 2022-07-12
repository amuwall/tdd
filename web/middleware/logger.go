package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		urlPath := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			urlPath = urlPath + "?" + c.Request.URL.RawQuery
		}

		log.Printf(
			"%s %s %s %s %d %s %s\n",
			c.ClientIP(),
			c.Request.Method,
			urlPath,
			c.Request.Proto,
			c.Writer.Status(),
			latencyTime,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
