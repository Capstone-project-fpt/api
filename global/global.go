package global

import (
	"github.com/api/pkg/logger"
	"github.com/api/pkg/setting"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config                 setting.Config
	I18nBundle             *i18n.Bundle
	Localizer              *i18n.Localizer
	Logger                 *logger.LoggerZap
	Db                     *gorm.DB
	RDb                    *redis.Client
)
