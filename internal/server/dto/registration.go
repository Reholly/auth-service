package dto

import (
	"github.com/pkg/errors"
	"net/mail"
)

const (
	UsernameMinLength = 3
	UsernameMaxLength = 30
	PasswordMinLength = 5
	PasswordMaxLength = 50
)

type Registration struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}

func (dto Registration) Validate() error {
	if len(dto.Username) > UsernameMaxLength ||
		len(dto.Username) < UsernameMinLength ||
		dto.Username == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: Имя пользователя - обязательное поле, должно быть в длину от 3 до  30 символов ")
	}

	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: Почта - обязательное поле, некорретные данные ")
	}

	if len(dto.Password) < PasswordMinLength ||
		len(dto.Password) > PasswordMaxLength ||
		dto.Password == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: Пароль - обязательное поле, должно быть в длину от 5 до  30 символов ")
	}

	if dto.Password != dto.RepeatedPassword {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поля паролей должны совпадать ")
	}

	return nil
}
