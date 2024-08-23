package middleware

import (
	"net/http"

	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *gin.Error:
				response.ErrorResponse(c, e.Meta.(int), e.Error())
				return
			default:
				response.ErrorResponse(c, http.StatusInternalServerError, "Service Unavailable")
				return
			}
		}
	}
}
