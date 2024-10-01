package initialize

import (
	"github.com/api/global"
	"github.com/hibiken/asynq"
)

func InitAsynq() {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: global.Config.Redis.Addr,
	})

	global.AsyncQClient = client
}