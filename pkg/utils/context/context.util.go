package context_util

import (
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/types"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetUserContext(ctx *gin.Context) *types.UserContext {
	userContext, isExist := ctx.Get("UserContext")

	if !isExist {
		response.ErrorResponse(ctx, 400, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TokenInvalid,
		}))
		return nil
	}

	return userContext.(*types.UserContext)
}