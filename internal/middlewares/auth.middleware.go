package middlewares

import (
	"go-ecommerce-be/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(response.ErrorCodeUnauthorized, gin.H{
				"message": "Token is required",
			})
			c.Abort()
			return
		}

		if token != "Bearer token" {
			c.JSON(response.ErrorCodeInvalidToken, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
