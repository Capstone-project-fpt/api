package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/service"
	"github.com/go-ecommerce-backend-api/pkg/response"
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
