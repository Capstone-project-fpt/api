package middleware

import (
	"net/http"

	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
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
