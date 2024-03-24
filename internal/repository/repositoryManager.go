package repository

import (
	"auth-service/internal/domain/repository"
)

type RepositoryManager struct {
	repository.AccountRepository
	repository.ClaimRepository
}

func NewRepositoryManager(
	accountRepo repository.AccountRepository,
	accountClaimRepo repository.ClaimRepository,
) *RepositoryManager {
	return &RepositoryManager{
		AccountRepository: accountRepo,
		ClaimRepository:   accountClaimRepo,
	}
}
