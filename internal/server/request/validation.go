package request

import (
	"auth-service/internal/domain/rules"
	"auth-service/internal/server/request/dto"
	"github.com/pkg/errors"
	"net/mail"
)

const (
	UsernameError      = "Имя пользователя не должно содерать спец.символы, должно начинаться с буквы и быть в длину от 5 до 30 символов"
	PasswordError      = "Пароль должен содержать цифры или буквы и не содержать спец.символы и должен быть в длину от 5 до 50 символов."
	EmailError         = "Некорретная почта."
	PasswordEqualError = "Пароли должны совпадать."
)

func ValidateUsernameDto(dto dto.Username) error {
	if !rules.IsUsernameValid(dto.Username) {
		return errors.Wrap(ErrorBadCredentials, UsernameError)
	}
	return nil
}
func ValidateLogInDto(dto dto.LogIn) error {
	if dto.Username == "" {
		return errors.Wrap(ErrorBadCredentials, "Поле <Имя пользователя> обязательно.")
	}

	if dto.Password == "" {
		return errors.Wrap(ErrorBadCredentials, "Поле <Пароль> обязательно.")
	}

	return nil
}
func ValidateRegistrationDto(dto dto.Registration) error {
	if !rules.IsUsernameValid(dto.Username) {
		return errors.Wrap(ErrorBadCredentials, UsernameError)
	}

	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return errors.Wrap(ErrorBadCredentials, EmailError)
	}

	if !rules.IsPasswordValid(dto.Password) {
		return errors.Wrap(ErrorBadCredentials, PasswordError)
	}

	if !rules.ArePasswordsEqual(dto.Password, dto.RepeatedPassword) {
		return errors.Wrap(ErrorBadCredentials, PasswordEqualError)
	}

	return nil
}
