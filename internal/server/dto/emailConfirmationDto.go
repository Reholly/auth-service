package dto

type EmailConfirmDto struct {
	Code     string `json:"code"`
	Username string `json:"username"`
}
