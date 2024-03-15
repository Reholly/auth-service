package repositories

import (
	"auth-service/internal/domain/entities"
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
func (r *AccountRepository) GetByUsernameWithClaims(ctx context.Context, username string) (entities.Account, error) {
	sql := `select u.id, u.username, u.email, u.hashed_password 
				from account u 
				where u.username = $1 
				limit 1`

	var account entities.Account

	accountRows, err := r.db.Query(ctx, sql, username)
	if err != nil {
		return entities.Account{}, err
	}
	defer accountRows.Close()

	accounts, err := pgx.CollectRows(accountRows, pgx.RowToStructByNameLax[dbModels.Account])
	if err != nil {
		return entities.Account{}, err
	}

	if len(accounts) == 0 {
		return entities.Account{}, storage.NotFoundErr
	}

	account = accounts[0].MapToEntity()

	sql = `select c.id, c.title, c.value
				from claim c
				left join account_claim ac on ac.claim_id = c.id
				where ac.account_id =$1`

	claimRows, err := r.db.Query(ctx, sql, account.Id)
	if err != nil {
		return entities.Account{}, err
	}
	defer accountRows.Close()

	claims, err := pgx.CollectRows(claimRows, pgx.RowToStructByNameLax[dbModels.Claim])

	for _, claim := range claims {
		account.Claims = append(account.Claims, claim.MapToEntity())
	}

	return account, nil
}

func (r *AccountRepository) UpdateAccountCredentials(ctx context.Context, account entities.Account) error {
	sql := `update account set username = $1, email = $2, hashed_password = $3`

	_, err := r.db.Exec(ctx, sql, account.Username, account.Email, account.HashedPassword)
	return err
}

func (r *AccountRepository) AddClaimToAccount(ctx context.Context, account entities.Account, claim entities.Claim) error {
	sql := `insert into account_claim (account_id, claim_id) values ($1, $2)`

	_, err := r.db.Exec(ctx, sql, account.Id, claim.Id)

	return err
}

func (r *AccountRepository) RemoveClaimFromAccount(ctx context.Context, account entities.Account, claim entities.Claim) error {
	sql := `delete from account_claim as ac where ac.account_id = $1 and ac.claim_id = $2`

	_, err := r.db.Exec(ctx, sql, account.Id, claim.Id)

	return err
}

func (r *AccountRepository) DeleteAccount(ctx context.Context, account entities.Account) error {
	sql := `delete from account as a where a.id = $1`

	_, err := r.db.Exec(ctx, sql, account.Id)

	return err
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account entities.Account) error {
	sql := `insert into account(username, email, hashed_password) values ($1, $2, $3)`

	_, err := r.db.Exec(ctx, sql, account.Username, account.Email, account.HashedPassword)

	return err
}
