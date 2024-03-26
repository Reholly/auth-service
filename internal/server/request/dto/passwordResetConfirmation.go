package dto

import (
	"auth-service/internal/server/request"
	"github.com/pkg/errors"
)

type PasswordResetConfirmation struct {
	Code             string
	Password         string
	RepeatedPassword string
}

func (dto PasswordResetConfirmation) Validate() error {
	if dto.Code == "" {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: поле <code> пусто ")
	}

	if dto.Password == "" {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: поле <password> пусто ")
	}

	if dto.RepeatedPassword == "" {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: поле <repeated password> пусто ")
	}

	if dto.RepeatedPassword != dto.Password {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: поля паролей должны совпадать ")
	}

	return nil
}
