package service

import (
	"auth-service/internal/config"
	"auth-service/internal/domain/entities"
	"auth-service/internal/domain/repositories"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthService struct {
	accountRepository repositories.AccountRepositoryContract
	config            config.Config
}

func (as *AuthService) LogIn(ctx context.Context, username, password string) (string, error) {
	account, err := as.accountRepository.GetByUsernameWithClaims(ctx, username)

	if err != nil {
		return "", err
	}

	if account.HashedPassword != as.hashPassword(password) {
		return "", ErrorWrongPassword
	}
	token, err := as.createToken(username, account.Claims)

	if err != nil {
		return "", err
	}

	return token, nil
}
func (as *AuthService) RegisterAccount(account entities.Account) (entities.Account, error) {

}
func (as *AuthService) ResetPassword(username, password string) error {

}
func (as *AuthService) DeleteAccountById(id uint64) error {

}

func (as *AuthService) createToken(username string, claims []entities.Claim) (string, error) {
	payload := jwt.MapClaims{}
	payload["username"] = username
	payload["exp"] = time.Now().Add(time.Hour * time.Duration(as.config.TokenTimeToLiveInHours)).Unix()
	for _, claim := range claims {
		payload[claim.Title] = claim.Value
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(as.config.JwtSecret)

	if err != nil {
		return "", err
	}

	return token, err
}

func (as *AuthService) hashPassword(password string) string {
	encoder := sha256.New()
	encoder.Write([]byte(password + as.config.Salt))
	return hex.EncodeToString(encoder.Sum(nil))
}
