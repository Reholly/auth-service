package implementations

import (
	"auth-service/internal/domain"
	"auth-service/internal/domain/helpers"
	"auth-service/internal/storage/postgres/repositories"
	"context"
	"github.com/pkg/errors"
)

type AdminService struct {
	repository  *repositories.RepositoryManager
	mailService domain.MailService
}

func NewAdminService(repository *repositories.RepositoryManager, mailService domain.MailService) domain.AdminService {
	return &AdminService{
		repository:  repository,
		mailService: mailService,
	}
}

func (s *AdminService) BanUser(ctx context.Context, username, reason string) error {
	err := s.repository.AccountRepository.UpdateAccountBanStatus(ctx, username, true)
	if err != nil {
		return errors.Wrap(err, "[ AdminService ] error user ban.")
	}

	account, err := s.repository.GetAccountByUsername(ctx, username)

	if err != nil {
		return errors.Wrap(err, "[ AdminService ] account repository error.")
	}

	err = s.mailService.SendMail(ctx, account.Email, "Бан", "Вы были забанены на платформе KForge.")

	return err
}

func (s *AdminService) UnbanUser(ctx context.Context, username string) error {
	err := s.repository.AccountRepository.UpdateAccountBanStatus(ctx, username, false)
	if err != nil {
		return errors.Wrap(err, "[ AdminService ] error user ban.")
	}

	account, err := s.repository.GetAccountByUsername(ctx, username)

	if err != nil {
		return err
	}

	err = s.mailService.SendMail(ctx, account.Email, "Разбан", "Вы были разбанены на платформе KForge.")

	return err
}
func (s *AdminService) CreateModerator(ctx context.Context, username string) error {
	err := s.repository.ClaimRepository.AddClaimByUsername(ctx, username, helpers.ModeratorRole)

	return errors.Wrap(err, "[ AdminService] error creating moderator.")
}

func (s *AdminService) DeleteModerator(ctx context.Context, username string) error {
	err := s.repository.ClaimRepository.RemoveClaimByUsername(ctx, username, helpers.ModeratorRole)

	return errors.Wrap(err, "[ AdminService] error deleting moderator.")
}
