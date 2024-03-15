package repositories

import (
	"auth-service/internal/domain/entities"
	"auth-service/internal/domain/repositories"
	"auth-service/internal/storage"
	"context"
	"github.com/jackc/pgx/v5"
)

type AccountRepository struct {
	db *pgx.Conn
}

func NewAccountRepository(db *pgx.Conn) *repositories.AccountRepositoryContract {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetByUsername(ctx context.Context, username string) (entities.Account, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password from account u where u.username = $1 limit 1`

	rows, err := r.db.Query(ctx, sql, username)
	var accounts []entities.Account
	if err != nil || len(accounts) == 0 {
		return entities.Account{}, storage.NotFoundErr
	}

	return url[0], nil
}

func (r *AccountRepository) UpdateAccount(account entities.Account) (entities.Account, error) {

}

func (r *AccountRepository) DeleteAccount(account entities.Account) error {

}

func (r *AccountRepository) CreateAccount(account entities.Account) (entities.Account, error) {

}
