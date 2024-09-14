package auth_service

import (
	"errors"
	"fmt"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (as *authService) LoginGoogleHandle(ctx *gin.Context) {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (as *authService) LoginGoogleCallbackHandle(ctx *gin.Context) (string, error) {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	userGoth, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintln(ctx.Writer, err)
		return "", err
	}

	var user model.User
	err = global.Db.Model(model.User{}).Select("id", "email", "password", "user_type", "name").Find(&user, "email = ?", userGoth.Email).Error

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return "", errors.New(message)
	}

	userContext := types.NewUserContext(&user)

	accessToken, refreshToken, err := as.authProcessService.ResolveAccessAndRefreshToken(ctx, &userContext)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return "", errors.New(message)
	}

	redirectUrl := fmt.Sprintf("%v?access_token=%v&refresh_token=%v", global.Config.Server.WebURL, accessToken, refreshToken)

	return redirectUrl, nil
}
