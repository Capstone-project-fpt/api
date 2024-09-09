package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/api/global"
	"go.uber.org/zap"
)

var (
	newAccountTemplate    = "new-account.html"
	resetPasswordTemplate = "reset-password.html"
)

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

type MailNewAccountTemplateData struct {
	Name     string
	Email    string
	Password string
}

type MailResetPasswordTemplateData struct {
	Name      string
	ResetLink string
}

func SendNewAccountEmail(to string, dataTemplate MailNewAccountTemplateData) error {
	htmlBody, err := getEmailTemplate(newAccountTemplate, dataTemplate)
	if err != nil {
		return err
	}

	return Send([]string{to}, global.Config.Smtp.Sender, "New account", htmlBody)
}

func SendResetPasswordEmail(to string, dataTemplate MailResetPasswordTemplateData) error {
	htmlBody, err := getEmailTemplate(resetPasswordTemplate, dataTemplate)
	if err != nil {
		return err
	}

	if global.Config.Server.Mode == "dev" {
		fmt.Println("Data reset password", dataTemplate)
	}

	return Send([]string{to}, global.Config.Smtp.Sender, "Reset password", htmlBody)
}

func Send(to []string, from string, subject string, htmlTemplate string) error {
	contentMail := Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    htmlTemplate,
	}

	message := BuildMessage(contentMail)

	smtpConfig := global.Config.Smtp

	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	if global.Config.Server.Mode != "dev" {
		err := smtp.SendMail(smtpConfig.Host+":"+smtpConfig.Port, auth, from, to, []byte(message))
		if err != nil {
			global.Logger.Error("Send mail error", zap.Error(err))
			return err
		}
	}
	fmt.Println("Send email Successful")

	return nil
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\r\n"
	msg += "Content-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ","))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += "\r\n"
	msg += fmt.Sprintf("%s\r\n", mail.Body)

	return msg
}

func getEmailTemplate(nameTemplate string, dataTemplate interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}
