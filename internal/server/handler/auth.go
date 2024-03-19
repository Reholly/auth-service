package handler

import (
	"auth-service/internal/domain"
	"auth-service/internal/server/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(service domain.AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var accountDto dto.AccountWithPasswordDto
	if err := c.ShouldBindJSON(&accountDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err := h.authService.RegisterAccount(c.Request.Context(), accountDto.Username, accountDto.Email, accountDto.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно зарегистрировать аккаунт. Возможно, аккаунт с таким username или email существует.")
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *AuthHandler) LogIn(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	token, err := h.authService.LogIn(c.Request.Context(), username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно войти. неправильный пароль или юзернейм.")
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *AuthHandler) ConfirmEmail(c *gin.Context) {
	code := c.Param("code")
	username := c.Param("username")

	err := h.authService.ConfirmEmail(c.Request.Context(), code, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Wrong code or username")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
func (h *AuthHandler) CreateModerator(c *gin.Context) {

}

func (h *AuthHandler) DeleteModerator(c *gin.Context) {

}

func (h *AuthHandler) BanUser(c *gin.Context) {

}

func (h *AuthHandler) UnbanUser(c *gin.Context) {

}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	username := c.Param("username")
	oldPassword := c.Param("oldPassword")
	newPassword := c.Param("newPassword")
	err := h.authService.ResetPassword(c.Request.Context(), username, oldPassword, newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно поменять пароль. Возможно, неправильно указан предыдущий пароль или же юзернейм.")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
