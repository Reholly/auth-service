package storage

import "errors"

var (
	ErrorNotFound      = errors.New("[Database] Not found")
	ErrorAlreadyExists = errors.New("[Database] Item already exists")
)
