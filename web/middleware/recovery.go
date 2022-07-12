package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				urlPath := c.Request.URL.Path
				if c.Request.URL.RawQuery != "" {
					urlPath = urlPath + "?" + c.Request.URL.RawQuery
				}

				log.Printf(
					"%s %s %s %s %s\n",
					c.ClientIP(), c.Request.Method, urlPath, c.Request.Proto, r,
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
