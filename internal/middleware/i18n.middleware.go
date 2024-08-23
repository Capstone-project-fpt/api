package middleware

import (
	"github.com/api/global"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bundle := global.I18nBundle

		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "vi"
		}

		localizer := i18n.NewLocalizer(bundle, lang)

		global.Localizer = localizer

		c.Next()
	}
}