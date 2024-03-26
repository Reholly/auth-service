package rules

import "regexp"

const (
	UsernameMinLength = 5
	UsernameMaxLength = 30
	PasswordMinLength = 5
	PasswordMaxLength = 50
)

var (
	//пароль должен состоять из символов обоих регистров и цифр
	PasswordRegex = regexp.MustCompile("^[0-9A-Za-z]+$")
	//имя пользователя не должно содержать спец.символы, а также не начинаться с цифры.
	UsernameRegex = regexp.MustCompile("^[^[:punct:]0-9]\\w*$")
)

func ArePasswordsEqual(password, repeatedPassword string) bool {
	return password == repeatedPassword
}

func IsPasswordValid(password string) bool {
	if len(password) > PasswordMaxLength || len(password) < PasswordMinLength {
		return false
	}

	return PasswordRegex.MatchString(password)
}

func IsUsernameValid(username string) bool {
	if len(username) > UsernameMaxLength || len(username) < UsernameMinLength {
		return false
	}

	return UsernameRegex.MatchString(username)
}
