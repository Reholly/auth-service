package interfaces

import (
	"auth-service/internal/domain"
	"context"
)

type ClaimRepository interface {
	GetClaimsByUsername(ctx context.Context, username string) ([]domain.Claim, error)
	AddClaimByUsername(ctx context.Context, username string, claim domain.Claim) error
	RemoveClaimByUsername(ctx context.Context, username string, claim domain.Claim) error
}
