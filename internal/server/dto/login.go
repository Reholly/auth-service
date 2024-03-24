package dto

import "github.com/pkg/errors"

type LogIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto LogIn) Validate() error {
	if dto.Username == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поле <Имя пользователя>  пустое ")
	}
	if dto.Username == "" {
		return errors.Wrap(ErrorBadCredentials, "Ошибка: поле <Пароль>  пустое ")
	}

	return nil
}
