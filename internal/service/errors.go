package service

import "errors"

var (
	ErrorWrongPassword = errors.New("[Service: AuthService] WrongPassword")
)
