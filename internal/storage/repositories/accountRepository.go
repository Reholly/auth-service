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
func (r *AccountRepository) GetByUsernameWithClaims(ctx context.Context, username string) (domain.Account, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password 
				from account u 
				where u.username = $1`

	var account domain.Account

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

	account = accounts[0].MapToEntity()

	sql = `select c.id, c.title, c.value
				from claim c
				left join account_claim ac on ac.claim_id = c.id
				where ac.account_id =$1`

	claimRows, err := r.db.Query(ctx, sql, account.Id)
	if err != nil {
		return domain.Account{}, err
	}
	defer accountRows.Close()

	claims, err := pgx.CollectRows(claimRows, pgx.RowToStructByNameLax[dbModels.Claim])

	for _, claim := range claims {
		account.Claims = append(account.Claims, claim.MapToEntity())
	}

	return account, nil
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
	sql := `update account as a set email = true where a.username = $1`

	_, err := r.db.Exec(ctx, sql, username)

	return err
}

func (r *AccountRepository) UpdateAccountCredentials(ctx context.Context, account domain.Account) error {
	sql := `update account set username = $1, email = $2, hashed_password = $3`

	_, err := r.db.Exec(ctx, sql, account.Username, account.Email, account.HashedPassword)

	return err
}

func (r *AccountRepository) AddClaimToAccount(ctx context.Context, account domain.Account, claim domain.Claim) error {
	sql := `insert into account_claim (account_id, claim_id) values ($1, $2)`

	_, err := r.db.Exec(ctx, sql, account.Id, claim.Id)

	return err
}

func (r *AccountRepository) RemoveClaimFromAccount(ctx context.Context, account domain.Account, claim domain.Claim) error {
	sql := `delete from account_claim as ac where ac.account_id = $1 and ac.claim_id = $2`

	_, err := r.db.Exec(ctx, sql, account.Id, claim.Id)

	return err
}

func (r *AccountRepository) DeleteAccountById(ctx context.Context, id uint64) error {
	sql := `delete from account as a where a.id = $1`

	_, err := r.db.Exec(ctx, sql, id)

	return err
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account domain.Account) error {
	sql := `insert into account(username, email, hashed_password) values ($1, $2, $3)`

	_, err := r.db.Exec(ctx, sql, account.Username, account.Email, account.HashedPassword)

	return err
}
