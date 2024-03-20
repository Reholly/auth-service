package repositories

import (
	"auth-service/internal/domain"
	"auth-service/internal/storage"
	"auth-service/internal/storage/postgres/models"
	"auth-service/lib/db"
	"context"
)

type AccountRepository struct {
	db *db.PostgresAdapter
}

func NewAccountRepository(db *db.PostgresAdapter) *AccountRepository {
	return &AccountRepository{db: db}
}
func (r *AccountRepository) GetAccountByUsername(ctx context.Context, username string) (domain.Account, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password, u.is_email_confirmed
				from account u 
				where u.username = $1`

	var accounts []models.Account
	err := r.db.Query(ctx, &accounts, sql, username)

	if len(accounts) == 0 {
		return domain.Account{}, storage.ErrorNotFound
	}

	return accounts[0].MapToEntity(), err
}

func (r *AccountRepository) CheckIfExistsAccountWithCredentials(ctx context.Context, username, email string) (bool, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password 
				from account u 
				where u.username = $1
				  and u.email = $2`

	var accounts []models.Account
	err := r.db.Query(ctx, &accounts, sql, username, email)

	if len(accounts) > 0 {
		return false, storage.ErrorAlreadyExists
	}

	return true, err
}

func (r *AccountRepository) ConfirmAccountEmail(ctx context.Context, username string) error {
	sql := `update account as a set is_email_confirmed = true where a.username = $1`

	return r.db.Execute(ctx, sql, username)
}

func (r *AccountRepository) UpdateAccountPassword(ctx context.Context, username, password string) error {
	sql := `update account as a set hashed_password = $1 where a.username = $2 `

	return r.db.Execute(ctx, sql, password, username)
}

func (r *AccountRepository) DeleteAccountById(ctx context.Context, id uint64) error {
	sql := `delete from account as a where a.id = $1`

	return r.db.Execute(ctx, sql, id)

}

func (r *AccountRepository) CreateAccount(ctx context.Context, username, email, hashedPassword string) error {
	sql := `insert into account(username, email, hashed_password) values ($1, $2, $3)`

	return r.db.Execute(ctx, sql, username, email, hashedPassword)
}
