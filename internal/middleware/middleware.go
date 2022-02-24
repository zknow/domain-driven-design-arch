package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// middleware logic
		c.Next()
	}
}
