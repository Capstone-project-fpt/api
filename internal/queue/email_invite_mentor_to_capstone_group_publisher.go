package queue

import (
	"encoding/json"
	"time"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type InviteMentorToCapstoneGroupMessage struct {
	MentorID          int
	MentorEmail       string
	CapstoneGroupID   int
	CapstoneGroupName string
	Token             string
}

type EmailInviteMentorToCapstoneGroupPublisher struct {
	client *asynq.Client
}

func NewEmailInviteMentorToCapstoneGroupPublisher() IBasePublisher[InviteMentorToCapstoneGroupMessage] {
	return &EmailInviteMentorToCapstoneGroupPublisher{
		client: global.AsyncqClient,
	}
}

func (e *EmailInviteMentorToCapstoneGroupPublisher) SendMessage(message InviteMentorToCapstoneGroupMessage, delay int) error {
	payload, err := json.Marshal(message)
	if err != nil {
		global.Logger.Error("Failed to marshal payload to send email task: ", zap.Error(err))
		return err
	}

	task := asynq.NewTask(
		constant.SystemQueueTask.SendEmailInviteMentorToCapstoneGroup,
		payload,
		asynq.ProcessIn(time.Duration(delay)*time.Second),
	)

	info, err := global.AsyncqClient.Enqueue(task)

	if err != nil {
		global.Logger.Error("Failed to send email task: ", zap.Error(err))
		return err
	}

	global.Logger.Info("Enqueued SendEmailInviteMentorToCapstoneGroup task: ", zap.String("ID", info.ID))

	return nil
}
