package initialize

import (
	"fmt"

	"github.com/api/global"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func InitGoth() {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd
	gothic.Store = store

	callBackURL := fmt.Sprintf("%s/api/v1/auth/google/callback", global.Config.Server.ServerURL)

	goth.UseProviders(
		google.New(
			global.Config.GoogleSetting.ClientID, 
			global.Config.GoogleSetting.ClientSecret, 
			callBackURL, 
			"email", 
			"profile",
		),
	)
}
