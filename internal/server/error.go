package server

import (
	"auth-service/internal/domain/service"
	"auth-service/internal/server/dto"
	"errors"
)

type APIError struct {
	Code    int
	Message string
}

func NewAPIError(err error) APIError {
	var message string
	var code int

	switch {
	case errors.Is(err, service.ErrorInvalidToken):
		message = err.Error()
		code = 403
	case errors.Is(err, service.ErrorAccountAlreadyExists),
		errors.Is(err, service.ErrorWrongEmailConfirmation),
		errors.Is(err, service.ErrorWrongPassword),
		errors.Is(err, dto.ErrorBadCredentials):

		message = err.Error()
		code = 400
	default:
		message = "Внутренная ошибка сервера. Обратитесь к администратору."
		code = 500
	}

	return APIError{
		Message: message,
		Code:    code,
	}
}
