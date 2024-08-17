package global

import (
	"github.com/go-ecommerce-backend-api/pkg/logger"
	"github.com/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MDb *gorm.DB
	RDb *redis.Client
)