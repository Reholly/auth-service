package models

import "auth-service/internal/domain"

type AccountClaim struct {
	Id              uint64 `db:"id"`
	AccountUsername string `db:"username"`
	Title           string `db:"claim_title"`
	Value           string `db:"claim_value"`
}

func (ac *AccountClaim) MapToEntity() domain.Claim {
	return domain.Claim{
		Title: ac.Title,
		Value: ac.Value,
	}
}
