package repositories

import "auth-service/internal/domain"

type RepositoryManager struct {
	domain.AccountRepository
	domain.AccountClaimRepository
}

func NewRepositoryManager(accountRepo domain.AccountRepository, accountClaimRepo domain.AccountClaimRepository) *RepositoryManager {
	return &RepositoryManager{
		AccountRepository:      accountRepo,
		AccountClaimRepository: accountClaimRepo,
	}
}
