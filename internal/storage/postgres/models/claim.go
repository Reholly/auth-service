package models

import (
	"auth-service/internal/domain/entity"
)

type AccountClaim struct {
	Id              uint64 `db:"id"`
	AccountUsername string `db:"username"`
	Title           string `db:"claim_title"`
	Value           string `db:"claim_value"`
}

func (ac *AccountClaim) MapToEntity() entity.Claim {
	return entity.Claim{
		Title: ac.Title,
		Value: ac.Value,
	}
}
