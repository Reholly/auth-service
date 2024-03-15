package repositories

import (
	"auth-service/internal/domain/entities"
	"context"
)

type AccountRepositoryContract interface {
	GetByUsernameWithClaims(ctx context.Context, username string) (entities.Account, error)
	UpdateAccountCredentials(ctx context.Context, account entities.Account) error
	AddClaimToAccount(ctx context.Context, account entities.Account, claim entities.Claim) error
	RemoveClaimFromAccount(ctx context.Context, account entities.Account, claim entities.Claim) error
	DeleteAccount(ctx context.Context, account entities.Account) error
	CreateAccount(ctx context.Context, account entities.Account) error
}
