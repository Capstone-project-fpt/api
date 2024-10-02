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

func HandleTaskEmailNewAccount(ctx context.Context, task *asynq.Task) error {
	var payload queue.EmailNewAccountsMessage
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		global.Logger.Error("Failed to unmarshal payload to EmailNewAccountsInput: ", zap.Error(err))
		return err
	}

	newAccounts := payload.NewAccounts
	for _, newAccount := range newAccounts {
		data := mail.MailNewAccountTemplateData{
			Name:     newAccount.Name,
			Email:    newAccount.Email,
			Password: newAccount.Password,
		}

		err := mail.SendNewAccountEmail(newAccount.Email, data)
		if err != nil {
			global.Logger.Error("Failed to send email to user with email: ", zap.String("email", newAccount.Email), zap.Error(err))
			continue
		}
		global.Logger.Info(fmt.Sprintf("Successfully sent email to %s", newAccount.Email))
	}

	return nil
}
