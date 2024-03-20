package handler

import (
	"auth-service/internal/server/dto"
	"auth-service/internal/service"
	"auth-service/internal/storage/postgres/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service    *service.ServiceManager
	repository *repositories.RepositoryManager
}

func NewAuthHandler(service *service.ServiceManager, repository *repositories.RepositoryManager) *AuthHandler {
	return &AuthHandler{
		service:    service,
		repository: repository,
	}
}

// @BasePath /api/v1

// Register endpoint godoc
// @Summary Endpoint for account registration.
// @Schemes
// @Description Account registration
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var accountDto dto.AccountWithPasswordDto
	if err := c.ShouldBindJSON(&accountDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err := h.service.RegisterAccount(c.Request.Context(), accountDto.Username, accountDto.Email, accountDto.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно зарегистрировать аккаунт. Возможно, аккаунт с таким username или email существует.")
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *AuthHandler) LogIn(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	token, err := h.service.LogIn(c.Request.Context(), username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно войти. неправильный пароль или юзернейм.")
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *AuthHandler) ConfirmEmail(c *gin.Context) {
	code := c.Param("code")
	username := c.Param("username")

	err := h.service.ConfirmEmail(c.Request.Context(), code, username)
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
	err := h.service.ResetPassword(c.Request.Context(), username, oldPassword, newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно поменять пароль. Возможно, неправильно указан предыдущий пароль или же юзернейм.")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
