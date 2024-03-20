package service

import (
	"auth-service/internal/config"
	"auth-service/internal/domain"
	"context"
	"fmt"
	"net/smtp"
)

type MailService struct {
	config *config.Config
}

func NewMailService(config *config.Config) domain.MailService {
	return &MailService{config: config}
}

func (ms *MailService) SendMail(ctx context.Context, address, header, message string) error {
	auth := smtp.PlainAuth("", ms.config.SmtpFrom, ms.config.SmtpPassword, ms.config.SmtpHost)
	finalMessage := fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n",
		address,
		header,
		message)

	err := smtp.SendMail(ms.config.SmtpHost+":"+ms.config.SmtpPort,
		auth,
		ms.config.SmtpFrom,
		[]string{
			address,
		},
		[]byte(finalMessage))

	if err != nil {
		return err
	}

	return nil
}
