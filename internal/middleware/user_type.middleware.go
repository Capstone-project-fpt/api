package middleware

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/pkg/response"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

func UserTypeMiddleware(allowUserTypes ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userContext := context_util.GetUserContext(ctx)

		if userContext == nil {
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.TokenInvalid,
			})
			response.ErrorResponse(ctx, http.StatusUnauthorized, message)
			ctx.Abort()
			return
		}

		if len(allowUserTypes) > 0 && !funk.Contains(allowUserTypes, userContext.UserType) {
			response.ErrorResponse(ctx, http.StatusForbidden, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.PermissionDenied,
			}))
			ctx.Abort()
			return
		}
		
		ctx.Next()
	}
}
