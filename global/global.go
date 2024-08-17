package global

import (
	"github.com/api/pkg/logger"
	"github.com/api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MDb    *gorm.DB
	RDb    *redis.Client
)
