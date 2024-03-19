package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/storage/repositories"
	"context"
)

type AdminService struct {
	repository *repositories.RepositoryManager
}

func NewAdminService(repository *repositories.RepositoryManager) domain.AdminService {
	return &AdminService{repository: repository}
}

func (s *AdminService) BanUser(ctx context.Context, username, reason string) error {
	return nil
}
func (s *AdminService) UnbanUser(ctx context.Context, username string) error {
	return nil
}
func (s *AdminService) CreateModerator(ctx context.Context, username string) error {
	return nil
}
func (s *AdminService) DeleteModerator(ctx context.Context, username string) error {
	return nil
}
