package repository

import (
	"auth-service/internal/domain/repository"
)

type RepositoryManager struct {
	repository.AccountRepository
	repository.ClaimRepository
	repository.CodeRepository
}

func NewRepositoryManager(
	accountRepo repository.AccountRepository,
	accountClaimRepo repository.ClaimRepository,
	codeRepository repository.CodeRepository,
) *RepositoryManager {
	return &RepositoryManager{
		AccountRepository: accountRepo,
		ClaimRepository:   accountClaimRepo,
		CodeRepository:    codeRepository,
	}
}
