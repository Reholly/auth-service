package repositories

import (
	"auth-service/internal/domain"
	"auth-service/internal/storage"
	"auth-service/internal/storage/dbModels"
	"context"
	"github.com/jackc/pgx/v5"
)

type AccountRepository struct {
	db *pgx.Conn
}

func NewAccountRepository(conn *pgx.Conn) *AccountRepository {
	return &AccountRepository{db: conn}
}
func (r *AccountRepository) GetAccountByUsername(ctx context.Context, username string) (domain.Account, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password, u.is_email_confirmed
				from account u 
				where u.username = $1`

	accountRows, err := r.db.Query(ctx, sql, username)
	if err != nil {
		return domain.Account{}, err
	}
	defer accountRows.Close()

	accounts, err := pgx.CollectRows(accountRows, pgx.RowToStructByNameLax[dbModels.Account])
	if err != nil {
		return domain.Account{}, err
	}

	if len(accounts) == 0 {
		return domain.Account{}, storage.ErrorNotFound
	}

	return accounts[0].MapToEntity(), nil
}

func (r *AccountRepository) CheckIfExistsAccountWithCredentials(ctx context.Context, username, email string) (bool, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password 
				from account u 
				where u.username = $1
				  and u.email = $2`

	accountRows, err := r.db.Query(ctx, sql, username, email)
	if err != nil {
		return false, err
	}
	defer accountRows.Close()

	accounts, err := pgx.CollectRows(accountRows, pgx.RowToStructByNameLax[dbModels.Account])
	if err != nil {
		return false, err
	}

	if len(accounts) > 0 {
		return false, storage.ErrorAlreadyExists
	}

	return true, nil
}

func (r *AccountRepository) ConfirmAccountEmail(ctx context.Context, username string) error {
	sql := `update account as a set is_email_confirmed = true where a.username = $1`

	_, err := r.db.Exec(ctx, sql, username)

	return err
}

func (r *AccountRepository) UpdateAccountPassword(ctx context.Context, username, password string) error {
	sql := `update account as a set hashed_password = $1 where a.username = $2 `

	_, err := r.db.Exec(ctx, sql, password, username)

	return err
}

func (r *AccountRepository) DeleteAccountById(ctx context.Context, id uint64) error {
	sql := `delete from account as a where a.id = $1`

	_, err := r.db.Exec(ctx, sql, id)

	return err
}

func (r *AccountRepository) CreateAccount(ctx context.Context, username, email, hashedPassword string) error {
	sql := `insert into account(username, email, hashed_password) values ($1, $2, $3)`

	_, err := r.db.Exec(ctx, sql, username, email, hashedPassword)

	return err
}
