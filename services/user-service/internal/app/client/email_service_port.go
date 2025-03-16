package client

import "github.com/zelalem-12/onetab/internal/app/service"

type EmailServicePort interface {
	SendEmail(sender string, receiver, cc, bcc []string, subject, templateName string, templateData *service.EmailContentData) error
}