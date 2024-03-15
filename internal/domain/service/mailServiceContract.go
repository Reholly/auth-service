package service

type MailServiceContract interface {
	ConfirmEmail(emailConfirmationCode string, email string) error
	SendEmailConfirmationMail(email string) error
}
