package dbModels

import "auth-service/internal/domain"

type Account struct {
	Id               uint64 `db:"id"`
	Username         string `db:"username"`
	Email            string `db:"email"`
	IsEmailConfirmed bool   `db:"email"`
	HashedPassword   string `db:"hashed_password"`
}

func (a *Account) MapToEntity() domain.Account {
	return domain.Account{
		Id:               a.Id,
		Username:         a.Username,
		Email:            a.Email,
		IsEmailConfirmed: a.IsEmailConfirmed,
		HashedPassword:   a.HashedPassword,
	}
}
