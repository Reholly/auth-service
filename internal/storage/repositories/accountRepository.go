package repositories

import (
	"auth-service/internal/domain/entities"
	"auth-service/internal/domain/repositories"
	"auth-service/internal/storage"
	"auth-service/internal/storage/dbModels"
	"context"
	"github.com/jackc/pgx/v5"
)

type AccountRepository struct {
	db *pgx.Conn
}

func NewAccountRepository(db *pgx.Conn) *repositories.AccountRepositoryContract {
	return &AccountRepository{db: db}
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

func (r *AccountRepository) UpdateAccount(account entities.Account) (entities.Account, error) {

}

func (r *AccountRepository) DeleteAccount(account entities.Account) error {

}

func (r *AccountRepository) CreateAccount(account entities.Account) (entities.Account, error) {

}
