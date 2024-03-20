package repositories

import (
	"auth-service/internal/storage/postgres/interfaces"
)

type RepositoryManager struct {
	interfaces.AccountRepository
	interfaces.ClaimRepository
}

func NewRepositoryManager(accountRepo interfaces.AccountRepository, accountClaimRepo interfaces.ClaimRepository) *RepositoryManager {
	return &RepositoryManager{
		AccountRepository: accountRepo,
		ClaimRepository:   accountClaimRepo,
	}
}
