package service

import "context"

type AuthService interface {
	LogIn(ctx context.Context, username, password string) (string, error)
	RegisterAccount(ctx context.Context, username, email, password string) error
	ResetPassword(ctx context.Context, code, newPassword string) error
	GenerateResetPasswordCode(ctx context.Context, username string) (string, error)
	ConfirmAccountEmail(ctx context.Context, code, username string) error
	DeleteAccountById(ctx context.Context, id uint64) error
}
