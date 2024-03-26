package dto

import (
	"auth-service/internal/server/request"
	"github.com/pkg/errors"
	"net/mail"
)

type PasswordResetRequest struct {
	Email string `json:"email"`
}

func (dto PasswordResetRequest) Validate() error {
	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return errors.Wrap(request.ErrorBadCredentials, "Ошибка: Почта - обязательное поле, некорретные данные ")
	}

	return nil
}
