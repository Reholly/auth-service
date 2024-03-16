package domain

import "context"

type AccountRepository interface {
	GetByUsernameWithClaims(ctx context.Context, username string) (Account, error)
	CheckIfExistsAccountWithCredentials(ctx context.Context, username, email string) (bool, error)
	UpdateAccountCredentials(ctx context.Context, account Account) error
	DeleteAccountById(ctx context.Context, id uint64) error
	CreateAccount(ctx context.Context, account Account) error
	ConfirmAccountEmail(ctx context.Context, username string) error

	AddClaimToAccount(ctx context.Context, account Account, claim Claim) error
	RemoveClaimFromAccount(ctx context.Context, account Account, claim Claim) error
}
