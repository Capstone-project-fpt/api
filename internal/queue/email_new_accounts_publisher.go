package queue

import (
	"encoding/json"
	"time"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type NewAccountMessage struct {
	Name     string
	Email    string
	Password string
}

type EmailNewAccountsMessage struct {
	NewAccounts []NewAccountMessage
}

type EmailNewAccountsPublisher struct {
	client *asynq.Client
}

func NewEmailNewAccountsPublisher() IBasePublisher[EmailNewAccountsMessage] {
	return &EmailNewAccountsPublisher{
		client: global.AsyncQClient,
	}
}

func (e *EmailNewAccountsPublisher) SendMessage(newAccounts EmailNewAccountsMessage, delay int) error {
	payload, err := json.Marshal(newAccounts)
	if err != nil {
		global.Logger.Error("Failed to marshal payload to send email task: ", zap.Error(err))
		return err
	}

	task := asynq.NewTask(
		constant.SystemQueueTask.SendEmailCreateAccounts,
		payload,
		asynq.ProcessIn(time.Duration(delay)*time.Second),
	)

	info, err := global.AsyncQClient.Enqueue(task)

	if err != nil {
		global.Logger.Error("Failed to send email task: ", zap.Error(err))
		return err
	}
	global.Logger.Info("Enqueued SendEmailCreateAccounts task: ", zap.String("ID", info.ID))

	return nil
}
