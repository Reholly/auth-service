package repository

import (
	"auth-service/internal/domain/entity"
	"context"
)

type AccountRepository interface {
	GetAccountByUsername(ctx context.Context, username string) (entity.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
	CheckIfAccountBanned(ctx context.Context, username string) (bool, error)
	CheckIfExistsAccountWithCredentials(ctx context.Context, username, email string) (bool, error)
	ConfirmAccountEmail(ctx context.Context, username string) error
	UpdateAccountBanStatus(ctx context.Context, username string, isBanned bool) error
	UpdateAccountPassword(ctx context.Context, username, password string) error
	DeleteAccountById(ctx context.Context, id uint64) error
	CreateAccount(ctx context.Context, username, email, hashedPassword string) error
}
