package dbModels

import "auth-service/internal/domain/entities"

type Account struct {
	Id             uint64 `db:"id"`
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
}

func (a Account) MapToEntity() entities.Account {
	return entities.Account{
		Id:             a.Id,
		Username:       a.Username,
		Email:          a.Email,
		HashedPassword: a.HashedPassword,
		Claims:         nil,
	}
}
