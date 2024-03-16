package service

import "errors"

var (
	ErrorWrongPassword          = errors.New("[Service: AuthService] Wrong password for account")
	ErrorAccountAlreadyExists   = errors.New("[Service: AuthService] Account already exists")
	ErrorWrongEmailConfirmation = errors.New("[Service: AuthService] Email confirmation error")
)
