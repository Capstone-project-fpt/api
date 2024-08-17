package initialize

import (
	"github.com/api/global"
	"github.com/api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(&global.Config.Logger)
}
