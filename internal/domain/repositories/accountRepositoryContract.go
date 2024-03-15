package repositories

import (
	"auth-service/internal/domain/entities"
	"context"
)

type AccountRepositoryContract interface {
	GetByUsernameWithClaims(ctx context.Context, username string) (entities.Account, error)
	UpdateAccount(ctx context.Context, account entities.Account) (entities.Account, error)
	DeleteAccount(ctx context.Context, account entities.Account) error
	CreateAccount(ctx context.Context, account entities.Account) (entities.Account, error)
}
