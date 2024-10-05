package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/auth_dto"
	auth_service "github.com/api/internal/service/auth"

	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AuthController struct {
	authService auth_service.IAuthService
}

func NewAuthController(authService auth_service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var input auth_dto.LoginInput
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

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body auth_dto.LoginInput true "data"
// @Router /login [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} auth_dto.LoginSwaggerOutput
func (ac *AuthController) Login(ctx *gin.Context) {
	var input auth_dto.LoginInput
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

	outputLogin := auth_dto.LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response.SuccessResponse(ctx, statusCode, outputLogin)
}

// @Summary Login With Google
// @Description Login With Google
// @Tags Auth
// @Produce json
// @Router /login/google [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
func (ac *AuthController) LoginGoogleHandle(ctx *gin.Context) {
	ac.authService.LoginGoogleHandle(ctx)

	response.SuccessResponse(
		ctx, http.StatusOK, "",
	)
}

func (ac *AuthController) LoginGoogleCallbackHandle(ctx *gin.Context) {
	redirectUrl, err := ac.authService.LoginGoogleCallbackHandle(ctx)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, redirectUrl)

	response.SuccessResponse(
		ctx, http.StatusOK, "",
	)
}

// @Summary ForgotPassword
// @Description Forgot Password
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body auth_dto.ForgotPasswordInput true "data"
// @Router /forgot-password [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
func (ac *AuthController) ForgotPassword(ctx *gin.Context) {
	var input auth_dto.ForgotPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	if err := ac.authService.ForgotPassword(ctx, input.Email); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "")
}

// @Summary ResetPassword
// @Description Reset Password
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body auth_dto.ResetPasswordInput true "data"
// @Router /reset-password [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
func (ac *AuthController) ResetPassword(ctx *gin.Context) {
	var input auth_dto.ResetPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	statusCode, err := ac.authService.ResetPassword(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, statusCode, "")
}
