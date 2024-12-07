package auth_service

import (
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

func (as *authService) LoginGoogleCallbackHandle(ctx *gin.Context) string {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	userGoth, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)

	url := fmt.Sprintf("%v/auth/sign-in", global.Config.Server.WebURL)

	if err != nil {
		fmt.Fprintln(ctx.Writer, err)
		url = fmt.Sprint(url, "?error=", err)

		return url
	}

	var user model.User
	err = global.Db.Model(model.User{}).Select("id", "email", "password", "user_type", "name").First(&user, "email = ?", userGoth.Email).Error

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		url = fmt.Sprint(url, "?error=", message)
		return url
	}

	userContext := types.NewUserContext(&user)

	accessToken, refreshToken, err := as.authProcessService.ResolveAccessAndRefreshToken(ctx, &userContext)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		url = fmt.Sprint(url, "?error=", message)

		return url
	}

	url = fmt.Sprintf("%v?access_token=%v&refresh_token=%v", url, accessToken, refreshToken)

	return url
}
