package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/auth_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var input auth_dto.InputLogin
	localizer := global.Localizer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	statusCode, err := ac.authService.Register(ctx, input.Email, input.Password)

	if err != nil {
		response.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

	response.SuccessResponse(ctx, 200, nil)
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var input auth_dto.InputLogin
	localizer := global.Localizer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	accessToken, refreshToken, statusCode, err := ac.authService.Login(ctx, input.Email, input.Password)

	if err != nil {
		response.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

	outputLogin := auth_dto.OutputLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response.SuccessResponse(ctx, statusCode, outputLogin)
}
