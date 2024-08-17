package initialize

import (
	"context"

	"github.com/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	redisConfig := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,  
		PoolSize: redisConfig.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis connection failed", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Init Redis success")
	global.RDb = rdb
}
