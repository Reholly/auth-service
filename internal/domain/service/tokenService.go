package domain

import domain "auth-service/internal/domain/entity"

type TokenService interface {
	CreateToken(claims []domain.Claim) (string, error)
	ParseClaims(jwtToken string) ([]domain.Claim, error)
}
