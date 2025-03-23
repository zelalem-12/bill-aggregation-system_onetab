package client

import "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/service"

type EmailServicePort interface {
	SendEmail(sender string, receiver, cc, bcc []string, subject, templateName string, templateData *service.EmailContentData) error
}
