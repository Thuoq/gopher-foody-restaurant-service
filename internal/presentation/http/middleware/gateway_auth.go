package middleware

import (
	"github.com/gin-gonic/gin"
)

func GatewayAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-Id")

		if userID != "" {
			c.Set("public_user_id", userID)
		}

		c.Next()
	}
}
