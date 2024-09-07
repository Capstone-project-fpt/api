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
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		token := strings.Split(authorization, "Bearer ")[1]
		redis := global.RDb
		key := token
		localizer := global.Localizer
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TokenInvalid,
		})

		val, err := redis.Get(c, key).Result()
		if err != nil {
			
			response.ErrorResponse(c, http.StatusUnauthorized, message)
			c.Abort()
			return
		}

		var userContext types.UserContext
		err = json.Unmarshal([]byte(val), &userContext)

		if err != nil {
			response.ErrorResponse(c, http.StatusUnauthorized, message)
			c.Abort()
			return
		}		

		c.Next()
	}
}
