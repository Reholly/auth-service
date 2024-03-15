package repositories

import "auth-service/internal/domain/entities"

type ClaimRepositoryContract interface {
	GetByTitle(username string) (entities.Account, error)
}
