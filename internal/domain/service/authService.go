package service

import "context"

type AuthService interface {
	LogIn(ctx context.Context, username, password string) (string, error)
	RegisterAccount(ctx context.Context, username, email, password string) error
	ResetPassword(ctx context.Context, username, oldPassword, newPassword string) error
	ConfirmEmail(ctx context.Context, code, username string) error
	DeleteAccountById(ctx context.Context, id uint64) error
}
