package client

import (
	"github.com/zelalem-12/onetab/internal/infrastructure/config"
	"gopkg.in/gomail.v2"
)

func InitSMTPClient(config *config.Config) *gomail.Dialer {
	return gomail.NewDialer(config.SMTP_HOST, config.SMTP_PORT, config.SENDER_EMAIL, config.SENDER_PASSWORD)
}
