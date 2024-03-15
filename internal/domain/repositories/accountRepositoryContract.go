package repositories

import "auth-service/internal/domain/entities"

type AccountRepositoryContract interface {
	GetByUsername(username string) (entities.Account, error)
	UpdateAccount(account entities.Account) (entities.Account, error)
	DeleteAccount(account entities.Account) error
	CreateAccount(account entities.Account) (entities.Account, error)
}
