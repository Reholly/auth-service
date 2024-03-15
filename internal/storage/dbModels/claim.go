package dbModels

import "auth-service/internal/domain/entities"

type Claim struct {
	Id    uint64 `db:"id"`
	Title string `db:"title"`
	Value string `db:"value"`
}

func (c Claim) MapToEntity() entities.Claim {
	return entities.Claim{
		Id:    c.Id,
		Title: c.Title,
		Value: c.Value,
	}
}
