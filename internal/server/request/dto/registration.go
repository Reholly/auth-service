package dto

const (
	UsernameMinLength = 3
	UsernameMaxLength = 30
	PasswordMinLength = 5
	PasswordMaxLength = 50
)

type Registration struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}
