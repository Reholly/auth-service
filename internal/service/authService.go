package service

import (
	"auth-service/internal/domain/entities"
	"auth-service/internal/domain/repositories"
)

type AuthService struct {
	accountRepository repositories.AccountRepositoryContract
}

func (as *AuthService) LogIn(username, password string) (string, error) {

	return "sd", nil
}
func (as *AuthService) RegisterAccount(account entities.Account) (entities.Account, error) {

}
func (as *AuthService) ResetPassword(username, password string) error {

}
func (as *AuthService) DeleteAccountById(id uint64) error {

}

func (as *AuthService) createToken() {

}
