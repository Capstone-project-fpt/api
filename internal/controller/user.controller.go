package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/user_dto"
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

// @Summary GetUser
// @Description Get User
// @Tags User
// @Produce json
// @Param id query int true "id"
// @Router /users/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} user_dto.OutputGetUserSwagger
// @Security ApiKeyAuth
func (u *UserController) GetUser(ctx *gin.Context) {
	var input user_dto.InputGetUser
	if err := ctx.ShouldBindQuery(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	outputGetUser, err := u.userService.GetUser(ctx, int(input.ID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, outputGetUser)
}

// @Summary GetMe
// @Description Get Me
// @Tags User
// @Produce json
// @Router /users/me [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} user_dto.OutputGetUserSwagger
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
