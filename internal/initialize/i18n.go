package initialize

import (
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/api/global"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitI18n() {
	absPath, err := filepath.Abs("internal/locales/")
	if err != nil {
		log.Fatalf("Failed to determine absolute path: %v", err)
		panic(err)
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile(filepath.Join(absPath, "en.json"))
	bundle.MustLoadMessageFile(filepath.Join(absPath, "vi.json"))

	global.I18nBundle = bundle
}