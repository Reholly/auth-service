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

// Register endpoint godoc
// @Summary Endpoint for account registration.
// @Description Account registration
// @Accept json
// @Produce json
// @Param input body dto.AccountWithPasswordDto true "Register"
// @Success 200 {string} ok
// @Failure 400 {string} bad request
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

// LogIn endpoint godoc
// @Summary Endpoint for getting access to system.
// @Description Log In
// @Accept json
// @Produce json
// @Param input body dto.LogInDto true "Register"
// @Success 200 {string} ok
// @Failure 400 {string} bad request
// @Router /api/auth/login [post]
func (h *AuthHandler) LogIn(c *gin.Context) {
	var loginDto dto.LogInDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	token, err := h.service.LogIn(c.Request.Context(), loginDto.Username, loginDto.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Невозможно войти. неправильный пароль или юзернейм.")
		return
	}
	c.JSON(http.StatusOK, token)
}

// ConfirmEmail endpoint godoc
// @Summary Endpoint for email confirmation after registration.
// @Description Confirm email
// @Accept json
// @Produce json
// @Param input query dto.EmailConfirmDto true "Confirm"
// @Success 200 {string} ok
// @Failure 400 {string} bad request
// @Router /api/auth/confirm [get]
func (h *AuthHandler) ConfirmEmail(c *gin.Context) {
	code := c.Query("code")
	username := c.Query("username")

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
