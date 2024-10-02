package worker

import (
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/worker/task"
	"github.com/hibiken/asynq"
)

func InitWorker() error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: global.Config.Redis.Addr,
		},
		asynq.Config{
			Concurrency: global.Config.AsynqSetting.MaxConcurrentWorkers,
		},
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc(constant.SystemQueueTask.SendEmailCreateAccounts, task.HandleTaskEmailNewAccount)

	return srv.Run(mux)
}