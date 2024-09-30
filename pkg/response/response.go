package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseDataSuccess struct {
	Code    int         `json:"code"`
	Success bool        `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErr struct {
	Code    int         `json:"code"`
	Success bool        `json:"message"`
	Error   interface{} `json:"error"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseDataSuccess{
		Code:    code,
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message interface{}) {
	c.JSON(code, ResponseErr{
		Code:    code,
		Success: false,
		Error:   message,
	})
}
