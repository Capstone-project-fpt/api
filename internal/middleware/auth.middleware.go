package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "valid-token" {
			response.ErrorResponse(c, http.StatusUnauthorized, "")
			c.Abort()
			return
		}

		c.Next()
	}
}