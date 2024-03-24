package interfaces

import (
	"auth-service/internal/domain/entity"
	"context"
)

type ClaimRepository interface {
	GetClaimsByUsername(ctx context.Context, username string) ([]entity.Claim, error)
	AddClaimByUsername(ctx context.Context, username string, claim entity.Claim) error
	RemoveClaimByUsername(ctx context.Context, username string, claim entity.Claim) error
}
