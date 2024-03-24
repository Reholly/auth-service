package implementation

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"

	"auth-service/internal/repository/models"
	"auth-service/lib/db"
	"context"
)

type AccountRepository struct {
	db *db.PostgresAdapter
}

func NewAccountRepository(db *db.PostgresAdapter) repository.AccountRepository {
	return &AccountRepository{db: db}
}
func (r *AccountRepository) GetAccountByUsername(ctx context.Context, username string) (entity.Account, error) {
	sql := `select u.id, u.username, u.is_banned, u.email, u.hashed_password, u.is_email_confirmed
				from account u 
				where u.username = $1`

	var accounts []models.Account
	err := r.db.Query(ctx, &accounts, sql, username)

	if len(accounts) == 0 {
		return entity.Account{}, repository.ErrorNotFound
	}

	return accounts[0].MapToEntity(), err
}

func (r *AccountRepository) CheckIfAccountBanned(ctx context.Context, username string) (bool, error) {
	sql := `select u.is_banned from account as u
				where u.username = $1
				  and u.email = $2`

	var accounts []models.Account
	err := r.db.Query(ctx, &accounts, sql, username)

	if len(accounts) == 0 {
		return false, repository.ErrorNotFound
	}

	return accounts[0].IsBanned, err
}

func (r *AccountRepository) CheckIfExistsAccountWithCredentials(ctx context.Context, username, email string) (bool, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password, u.is_banned
				from account as u 
				where u.username = $1
				  and u.email = $2`

	var accounts []models.Account
	err := r.db.Query(ctx, &accounts, sql, username, email)

	if len(accounts) > 0 {
		return true, repository.ErrorAlreadyExists
	}

	return false, err
}

func (r *AccountRepository) ConfirmAccountEmail(ctx context.Context, username string) error {
	sql := `update account as a set is_email_confirmed = true where a.username = $1`

	return r.db.Execute(ctx, sql, username)
}

func (r *AccountRepository) UpdateAccountPassword(ctx context.Context, username, password string) error {
	sql := `update account as a set hashed_password = $1 where a.username = $2 `

	return r.db.Execute(ctx, sql, password, username)
}

func (r *AccountRepository) UpdateAccountBanStatus(ctx context.Context, username string, isBanned bool) error {
	sql := `update account as a set is_banned = $1 where a.username = $2 `

	return r.db.Execute(ctx, sql, isBanned, username)
}

func (r *AccountRepository) DeleteAccountById(ctx context.Context, id uint64) error {
	sql := `delete from account as a where a.id = $1`

	return r.db.Execute(ctx, sql, id)

}

func (r *AccountRepository) CreateAccount(ctx context.Context, username, email, hashedPassword string) error {
	sql := `insert into account(username, email, hashed_password, is_banned, is_email_confirmed) 
				values ($1, $2, $3, false, false)`

	return r.db.Execute(ctx, sql, username, email, hashedPassword)
}
