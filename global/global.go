package global

import (
	database "github.com/api/database/sqlc"
	"github.com/api/pkg/logger"
	"github.com/api/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Db     *database.Queries
	RDb    *redis.Client
)
