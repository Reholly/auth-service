package service

import "auth-service/internal/domain/entity"

type TokenService interface {
	CreateToken(claims []entity.Claim) (string, error)
	ParseClaims(jwtToken string) ([]entity.Claim, error)
	IsModerator(claims []entity.Claim) bool
	IsAdmin(claims []entity.Claim) bool
	IsInAdministration(claims []entity.Claim) bool
}
