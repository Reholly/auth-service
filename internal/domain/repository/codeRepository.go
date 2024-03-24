package repository

type CodeType string

var (
	EmailConfirmation CodeType = "email_confirm"
	PasswordReset     CodeType = "reset"
)

type CodeRepository interface {
	GetByUsername(username string, codeType CodeType) (string, error)
	GetUsernameByCode(code string) (string, error)
	Set(username, code string, codeType CodeType)
	Remove(username string, codeType CodeType)
}
