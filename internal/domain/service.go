package domain

import "context"

type AuthService interface {
	LogIn(ctx context.Context, username, password string) (string, error)
	RegisterAccount(ctx context.Context, username, email, password string) error
	ResetPassword(ctx context.Context, username, oldPassword, newPassword string) error
	ConfirmEmail(ctx context.Context, code, username string) error
	DeleteAccountById(ctx context.Context, id uint64) error
}

type MailService interface {
	SendMail(ctx context.Context, header, message string) error
}

type TokenService interface {
	CreateToken(claims []Claim) (string, error)
	ParseClaims(jwtToken string) ([]Claim, error)
}
