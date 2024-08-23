package controller

import (
	"net/http"

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

func (userController *UserController) Register(ctx *gin.Context) {
	_, err := userController.userService.Register(ctx, "email", "password")

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	response.SuccessResponse(ctx, http.StatusAccepted, nil)
}
