package client

import (
	"bytes"
	"fmt"
	"html/template"

	clientPort "github.com/zelalem-12/onetab/internal/app/client"
	"github.com/zelalem-12/onetab/internal/app/service"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	smtpClient  *gomail.Dialer
	templateDir string
}

func NewEmailService(smtpClient *gomail.Dialer) clientPort.EmailServicePort {
	return &EmailService{
		smtpClient:  smtpClient,
		templateDir: "internal/adapter/template",
	}
}

func (e *EmailService) SendEmail(sender string, receiver, cc, bcc []string, subject, templateName string, templateData *service.EmailContentData) error {
	if sender == "" {
		return fmt.Errorf("sender email cannot be empty")
	}

	if len(receiver) == 0 {
		return fmt.Errorf("receiver email cannot be empty")
	}

	if subject == "" {
		return fmt.Errorf("email subject cannot be empty")
	}

	if templateName == "" {
		return fmt.Errorf("template name cannot be empty")
	}

	if templateData == nil {
		return fmt.Errorf("template data cannot be nil")
	}

	content, err := e.renderTemplate(templateName, templateData)
	if err != nil {
		return fmt.Errorf("failed to render template: %v", err)
	}

	return e.smtpClient.DialAndSend(e.composeMailer(sender, receiver, cc, bcc, subject, content))
}

func (e *EmailService) renderTemplate(templateName string, data *service.EmailContentData) (string, error) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%s/base.html", e.templateDir),
		fmt.Sprintf("%s/%s.html", e.templateDir, templateName),
	)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (e *EmailService) composeMailer(sender string, receiver, cc, bcc []string, subject, content string) *gomail.Message {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", sender)
	mailer.SetHeader("To", receiver...)

	if len(cc) > 0 {
		mailer.SetHeader("Cc", cc...)
	}

	if len(bcc) > 0 {
		mailer.SetHeader("Bcc", bcc...)
	}

	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", content)

	return mailer
}
