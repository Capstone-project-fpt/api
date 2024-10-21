package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/api/global"
	"github.com/api/internal/queue"
	"github.com/api/pkg/mail"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func HandleTaskEmailInviteMentorToCapstoneGroup(ctx context.Context, task *asynq.Task) error {
	var payload queue.InviteMentorToCapstoneGroupMessage
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		global.Logger.Error("Failed to unmarshal payload to EmailInviteMentorToCapstoneGroupInput: ", zap.Error(err))
		return err
	}

	data := mail.MailInviteMentorToCapstoneGroupTemplateData{
		MentorID:          payload.MentorID,
		MentorEmail:       payload.MentorEmail,
		CapstoneGroupID:   payload.CapstoneGroupID,
		CapstoneGroupName: payload.CapstoneGroupName,
		Token:             payload.Token,
	}

	fmt.Println("data: ", data)

	err := mail.SendEmailInviteMentorToCapstoneGroup(data.MentorEmail, data)
	if err != nil {
		global.Logger.Error("Failed to send email to user with email: ", zap.String("email", data.MentorEmail), zap.Error(err))
		return err
	}
	global.Logger.Info(fmt.Sprintf("Successfully sent email to %s", data.MentorEmail))

	return nil
}
