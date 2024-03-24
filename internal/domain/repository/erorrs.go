package repository

import "errors"

var (
	ErrorNotFound      = errors.New("[ Repository ] Не найдено. ")
	ErrorAlreadyExists = errors.New("[ Repository ] Уже существует. ")
)
