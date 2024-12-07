package initialize

import (
	"fmt"

	"github.com/api/global"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func InitGoth() {
	callBackURL := fmt.Sprintf("%s/api/v1/auth/google/callback", global.Config.Server.ServerURL)

	clientID := global.Config.GoogleSetting.ClientID
	clientSecret := global.Config.GoogleSetting.ClientSecret

	goth.UseProviders(
		google.New(clientID, clientSecret, callBackURL),
	)

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	gothic.Store = store
}
