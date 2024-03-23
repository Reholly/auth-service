package models

import "auth-service/internal/domain"

type Account struct {
	Id               uint64 `db:"id"`
	Username         string `db:"username"`
	Email            string `db:"email"`
	IsEmailConfirmed bool   `db:"is_email_confirmed"`
	IsBanned         bool   `db:"is_banned"`
	HashedPassword   string `db:"hashed_password"`
}

func (a *Account) MapToEntity() domain.Account {
	return domain.Account{
		Id:               a.Id,
		Username:         a.Username,
		Email:            a.Email,
		IsBanned:         a.IsBanned,
		IsEmailConfirmed: a.IsEmailConfirmed,
		HashedPassword:   a.HashedPassword,
	}
}
