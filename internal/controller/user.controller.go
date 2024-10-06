package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// @Summary GetMe
// @Description Get Me
// @Tags User
// @Produce json
// @Router /users/me [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} user_dto.GetUserSwaggerOutput
// @Security ApiKeyAuth
func (u *UserController) GetMe(ctx *gin.Context) {
	userContext := context_util.GetUserContext(ctx)
	if userContext == nil {
		response.ErrorResponse(ctx, http.StatusNotFound, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
		return
	}

	outputGetUser, err := u.userService.GetUser(ctx, int(userContext.ID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, outputGetUser)
}
