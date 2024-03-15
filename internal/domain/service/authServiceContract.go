package service

import "auth-service/internal/domain/entities"

type AuthServiceContract interface {
	LogIn(username, password string) (string, error)
	RegisterAccount(account entities.Account) (entities.Account, error)
	ResetPassword(username, password string) error
	DeleteAccountById(id uint64) error
}
