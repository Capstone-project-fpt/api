package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/types"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

func AuthMiddleware(allowUserTypes ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		splitAuthorization := strings.Split(authorization, "Bearer ")
		if len(splitAuthorization) != 2 {
			response.ErrorResponse(ctx, http.StatusUnauthorized, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.TokenInvalid,
			}))
			ctx.Abort()
			return
		}

		token := splitAuthorization[1]
		redis := global.RDb
		key := token
		localizer := global.Localizer
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TokenInvalid,
		})

		val, err := redis.Get(ctx, key).Result()
		if err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, message)
			ctx.Abort()
			return
		}

		var userContext types.UserContext
		if err = json.Unmarshal([]byte(val), &userContext); err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, message)
			ctx.Abort()
			return
		}

		if len(allowUserTypes) > 0 && !funk.Contains(allowUserTypes, userContext.UserType) {
			message = localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.PermissionDenied,
			})

			response.ErrorResponse(ctx, http.StatusUnauthorized, message)
			ctx.Abort()
			return
		}

		ctx.Set("UserContext", &userContext)

		ctx.Next()
	}
}
