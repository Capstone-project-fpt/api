package middleware

import (
	"net/http"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/pkg/response"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func PermissionMiddleware(permission ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		localizer := global.Localizer
		userContext := context_util.GetUserContext(ctx)

		if userContext == nil {
			message := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.TokenInvalid,
			})
			response.ErrorResponse(ctx, http.StatusUnauthorized, message)
			ctx.Abort()
			return
		}

		var userPermissions []model.Permission
		err := global.Db.Raw(`
    SELECT id, name FROM permissions 
    WHERE id IN (
        SELECT permission_id FROM roles_permissions
        WHERE role_id IN (
            SELECT role_id FROM users_roles
            WHERE user_id = ?
        )
    )
		`, userContext.ID).Scan(&userPermissions).Error

		if err != nil {
			message := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.PermissionDenied,
			})
			response.ErrorResponse(ctx, http.StatusForbidden, message)
			ctx.Abort()
			return
		}

		isExist := false
		for _, permission := range permission {
			for _, userPermission := range userPermissions {
				if userPermission.Name == permission {
					isExist = true
					break
				}
			}
		}

		if !isExist {
			message := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.PermissionDenied,
			})
			response.ErrorResponse(ctx, http.StatusForbidden, message)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
