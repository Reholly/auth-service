package service

import (
	"auth-service/internal/config"
	"auth-service/internal/domain"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
)

type AuthService struct {
	accountRepository domain.AccountRepository
	tokenService      domain.TokenService
	mailService       domain.MailService
	config            config.Config
}

func (as *AuthService) LogIn(ctx context.Context, username, password string) (string, error) {
	account, err := as.accountRepository.GetByUsernameWithClaims(ctx, username)

	if err != nil {
		return "", err
	}

	if account.HashedPassword != as.hash(password) {
		return "", ErrorWrongPassword
	}
	account.Claims = append(account.Claims, domain.Claim{
		Title: "username",
		Value: account.Username,
	})

	token, err := as.tokenService.CreateToken(account.Claims)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) RegisterAccount(ctx context.Context, account domain.Account) error {
	ok, err := as.accountRepository.CheckIfExistsAccountWithCredentials(ctx, account.Username, account.Email)

	if err != nil {
		return err
	}

	if !ok {
		return ErrorAccountAlreadyExists
	}

	err = as.accountRepository.CreateAccount(ctx, account)
	if err != nil {
		return err
	}
	confirmationCode := as.hash(account.Username)

	params := url.Values{}
	params.Set("username", account.Username)
	params.Set("code", confirmationCode)
	confirmationUrl := fmt.Sprintf("%s?%s", as.config.EmailConfirmationUrlBase, params.Encode())
	confirmationMessage := fmt.Sprintf("Для подтверждения почты перейдите по ссылке: <a href='%s'>ссылке.</a>", confirmationUrl)
	err = as.mailService.SendMail(ctx, "Подтверждение почты", confirmationMessage)

	if err != nil {
		return err
	}

	as.accountRepository.CreateAccount(ctx, account)

	return nil
}

func (as *AuthService) ConfirmEmail(ctx context.Context, code, username string) error {
	if as.hash(username) != code {
		return ErrorWrongEmailConfirmation
	}
	return as.accountRepository.ConfirmAccountEmail(ctx, username)
}

func (as *AuthService) ResetPassword(ctx context.Context, username, oldPassword, newPassword string) error {
	account, err := as.accountRepository.GetByUsernameWithClaims(ctx, username)

	if err != nil {
		return err
	}

	if account.HashedPassword != as.hash(oldPassword) {
		return ErrorWrongPassword
	}
	account.HashedPassword = newPassword

	err = as.accountRepository.UpdateAccountCredentials(ctx, account)
	if err != nil {
		return err
	}
	return nil
}
func (as *AuthService) DeleteAccountById(ctx context.Context, id uint64) error {

	err := as.accountRepository.DeleteAccountById(ctx, id)
	return err
}

func (as *AuthService) hash(item string) string {
	encoder := sha256.New()
	encoder.Write([]byte(item + as.config.Salt))
	return hex.EncodeToString(encoder.Sum(nil))
}
