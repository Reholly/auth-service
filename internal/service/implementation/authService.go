package implementation

import (
	"auth-service/config"
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/helpers"
	"auth-service/internal/domain/service"
	"auth-service/internal/repository"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
)

type AuthService struct {
	repository   *repository.RepositoryManager
	mailService  service.MailService
	tokenService service.TokenService
	config       *config.Config
}

func NewAuthService(
	config *config.Config,
	repository *repository.RepositoryManager,
	mailService service.MailService,
	tokenService service.TokenService) service.AuthService {
	return &AuthService{
		repository:   repository,
		mailService:  mailService,
		tokenService: tokenService,
		config:       config,
	}
}

func (as *AuthService) LogIn(ctx context.Context, username, password string) (string, error) {
	dbAccount, err := as.repository.GetAccountByUsername(ctx, username)

	if err != nil {
		return "", err
	}

	if dbAccount.HashedPassword != as.hash(password) {
		return "", service.ErrorWrongPassword
	}

	claims, err := as.repository.GetClaimsByUsername(ctx, username)
	claims = append(claims, entity.Claim{
		Title: "username",
		Value: username,
	})

	if err != nil {
		return "", err
	}

	token, err := as.tokenService.CreateToken(claims)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) RegisterAccount(ctx context.Context, username, email, password string) error {
	exist, err := as.repository.CheckIfExistsAccountWithCredentials(ctx, username, email)

	if err != nil {
		return err
	}

	if exist {
		return service.ErrorAccountAlreadyExists
	}

	err = as.repository.CreateAccount(ctx, username, email, as.hash(password))

	if err != nil {
		return err
	}

	err = as.repository.AddClaimByUsername(ctx, username, helpers.StudentRole)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	confirmationCode := as.hash(username)

	params := url.Values{}
	params.Set("username", username)
	params.Set("code", confirmationCode)
	confirmationUrl := fmt.Sprintf("%s?%s", as.config.EmailConfirmationUrlBase, params.Encode())
	confirmationMessage := fmt.Sprintf("Для подтверждения почты перейдите по: <a href=\"%s\">ссылке.</a>", confirmationUrl)
	err = as.mailService.SendMail(ctx, email, "Подтверждение почты", confirmationMessage)

	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) ConfirmEmail(ctx context.Context, code, username string) error {
	if as.hash(username) != code {
		return service.ErrorWrongEmailConfirmation
	}
	return as.repository.ConfirmAccountEmail(ctx, username)
}

func (as *AuthService) ResetPassword(ctx context.Context, username, oldPassword, newPassword string) error {
	account, err := as.repository.GetAccountByUsername(ctx, username)

	if err != nil {
		return err
	}

	if account.HashedPassword != as.hash(oldPassword) {
		return service.ErrorWrongPassword
	}

	err = as.repository.UpdateAccountPassword(ctx, username, newPassword)
	if err != nil {
		return err
	}
	return nil
}
func (as *AuthService) DeleteAccountById(ctx context.Context, id uint64) error {
	err := as.repository.DeleteAccountById(ctx, id)
	return err
}

func (as *AuthService) hash(item string) string {
	encoder := sha256.New()
	encoder.Write([]byte(item + as.config.Salt))
	return hex.EncodeToString(encoder.Sum(nil))
}
