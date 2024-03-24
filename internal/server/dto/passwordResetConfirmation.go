package dto

import "github.com/pkg/errors"

type PasswordResetConfirmation struct {
	Code             string
	Password         string
	RepeatedPassword string
}

func (dto PasswordResetConfirmation) Validate() error {
	if dto.Code == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поле <code> пусто ")
	}

	if dto.Password == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поле <password> пусто ")
	}

	if dto.RepeatedPassword == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поле <repeated password> пусто ")
	}

	if dto.RepeatedPassword != dto.Password {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поля паролей должны совпадать ")
	}

	return nil
}
