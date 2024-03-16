package domain

import (
	"context"
)

type AccountRepository interface {
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
	CheckIfExistsAccountWithCredentials(ctx context.Context, username, email string) (bool, error)
	ConfirmAccountEmail(ctx context.Context, username string) error
	UpdateAccountPassword(ctx context.Context, username, password string) error
	DeleteAccountById(ctx context.Context, id uint64) error
	CreateAccount(ctx context.Context, username, email, hashedPassword string) error
}

type AccountClaimRepository interface {
	GetClaimsByUsername(ctx context.Context, username string) ([]Claim, error)
	AddClaimByUsername(ctx context.Context, username string, claim Claim) error
}
