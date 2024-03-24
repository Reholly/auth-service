package repository

import "errors"

var (
	ErrorNotFound      = errors.New("[ Repository ] Not found")
	ErrorAlreadyExists = errors.New("[ Repostiroy ] Item already exists")
)
