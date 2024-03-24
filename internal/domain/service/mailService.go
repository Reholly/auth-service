package domain

import "context"

type MailService interface {
	SendMail(ctx context.Context, address, header, message string) error
}
