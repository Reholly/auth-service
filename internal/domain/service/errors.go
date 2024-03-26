package service

import "errors"

var (
	ErrorWrongPassword            = errors.New("[ Service: AuthService ] Неправильный пароль. ")
	ErrorAccountAlreadyExists     = errors.New("[ Service: AuthService ] Аккаунт уже существует. ")
	ErrorWrongEmailConfirmation   = errors.New("[ Service: AuthService ] Ошибка подтверждения почты: неправильный код или имя пользователя. ")
	ErrorInvalidResetPasswordCode = errors.New("[ Service: AuthService ] Некорректный код сброса пароля. ")

	ErrorInvalidToken = errors.New("[ Service: TokenService ] Некорректный токен. ")

	ErrorInvalidPermissions = errors.New("[ Service: AdminService] Недостаточно прав для доступа. ")
)
