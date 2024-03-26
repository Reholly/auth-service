package dto

import (
	"auth-service/internal/server/request"
	"github.com/pkg/errors"
)

type EmailConfirmation struct {
	Code     string `json:"code"`
	Username string `json:"username"`
}

func (dto EmailConfirmation) Validate() error {
	if dto.Code == "" {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: параметр <code>  пуст ")
	}
	if dto.Username == "" {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: параметр <username>  пуст ")
	}

	return nil
}
