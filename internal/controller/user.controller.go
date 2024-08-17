package controller

import (
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) Register(c *gin.Context) {
	result := userController.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
