package repositories

import (
	"auth-service/internal/domain"
	"auth-service/internal/storage/dbModels"
	"context"
	"github.com/jackc/pgx/v5"
)

type AccountClaimRepository struct {
	db *pgx.Conn
}

func NewClaimRepository(conn *pgx.Conn) AccountClaimRepository {
	return AccountClaimRepository{db: conn}
}

func (r *AccountClaimRepository) GetClaimsByUsername(ctx context.Context, username string) ([]domain.Claim, error) {
	sql := `select ac.id, ac.username, ac.claim_title, ac.claim_value
				from account_claim as ac 
				where ac.username = $1`

	rows, err := r.db.Query(ctx, sql, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbClaims, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[dbModels.AccountClaim])

	if err != nil {
		return nil, err
	}

	claims := make([]domain.Claim, len(dbClaims))
	for _, dbClaim := range dbClaims {
		claims = append(claims, domain.Claim{
			Title: dbClaim.Title,
			Value: dbClaim.Value,
		})
	}

	return claims, nil
}

func (r *AccountClaimRepository) AddClaimByUsername(ctx context.Context, username string, claim domain.Claim) error {
	sql := `insert into account_claim(username, claim_title, claim_value) values ($1, $2, $3)`

	_, err := r.db.Exec(ctx, sql, username, claim.Title, claim.Value)

	return err
}
