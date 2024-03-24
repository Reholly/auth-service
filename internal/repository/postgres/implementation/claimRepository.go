package implementation

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"
	"auth-service/internal/repository/models"
	"auth-service/lib/db"
	"context"
)

type AccountClaimRepository struct {
	db *db.PostgresAdapter
}

func NewClaimRepository(conn *db.PostgresAdapter) repository.ClaimRepository {
	return &AccountClaimRepository{db: conn}
}

func (r *AccountClaimRepository) GetClaimsByUsername(ctx context.Context, username string) ([]entity.Claim, error) {
	sql := `select ac.id, ac.username, ac.claim_title, ac.claim_value
				from account_claim as ac 
				where ac.username = $1`

	var dbClaims []models.AccountClaim
	err := r.db.Query(ctx, &dbClaims, sql, username)

	if err != nil {
		return nil, err
	}

	claims := make([]entity.Claim, len(dbClaims))
	for _, dbClaim := range dbClaims {
		claims = append(claims, entity.Claim{
			Title: dbClaim.Title,
			Value: dbClaim.Value,
		})
	}

	return claims, nil
}

func (r *AccountClaimRepository) AddClaimByUsername(ctx context.Context, username string, claim entity.Claim) error {
	sql := `insert into account_claim(username, claim_title, claim_value) values ($1, $2, $3)`

	return r.db.Execute(ctx, sql, username, claim.Title, claim.Value)
}
func (r *AccountClaimRepository) RemoveClaimByUsername(ctx context.Context, username string, claim entity.Claim) error {
	sql := `delete from account_claim where username = $1 and claim_title= $2 and claim_value = $3`

	return r.db.Execute(ctx, sql, username, claim.Title, claim.Value)
}
