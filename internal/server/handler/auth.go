package handler

import (
	"auth-service/internal/repository"
	"auth-service/internal/server"
	"auth-service/internal/server/request/dto"
	"auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service    *service.ServiceManager
	repository *repository.RepositoryManager
}

func NewAuthHandler(service *service.ServiceManager, repository *repository.RepositoryManager) *AuthHandler {
	return &AuthHandler{
		service:    service,
		repository: repository,
	}
}

// Register godoc
// @Summary Endpoint for account registration.
// @Description Этот эндпоинт служит для выполнения регистрации новых аккаунтов на KForge.
// @Description С эндпоинта можно получить 200, 400, 500 статус-коды.
// @Description Ошибки с эндпоинта: Bad Credentials, Internal Server Error
// @Accept json
// @Produce json
// @Param input body request.Registration true "Register"
// @Success 200 {string} ok
// @Failure 400 {string} bad request: error in credentials
// @Failure 500 {string} server error: registration error
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var accountDto dto.Registration
	if err := c.ShouldBindJSON(&accountDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := accountDto.Validate(); err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	err := h.service.RegisterAccount(c.Request.Context(), accountDto.Username, accountDto.Email, accountDto.Password)

	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// LogIn godoc
// @Summary Endpoint for log in system.
// @Description Эндпоинт для входа на KForge. На выходе эндпоинт отдает JWT токен с набором следующих claim:
// @Description 1) role : student / admin / moderator
// @Description 2) exp : время, когда токен перестанет действовать.
// @Description 3) username : имя пользователя. Вообще, нужно для связи с остальными микросервисами.
// @Description Ошибки с эндпоинта: Bad Credentials, Internal Server Error
// @Accept json
// @Produce json
// @Param input body request.LogIn true "Log In"
// @Success 200 {string} ok
// @Failure 400 {string} bad request
// @Failure 500 {string} internal server error
// @Router /api/auth/login [post]
func (h *AuthHandler) LogIn(c *gin.Context) {
	var loginDto dto.LogIn
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := loginDto.Validate(); err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	token, err := h.service.LogIn(c.Request.Context(), loginDto.Username, loginDto.Password)

	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	c.JSON(http.StatusOK, token)
}

// ConfirmEmail godoc
// @Summary Endpoint for email confirmation.
// @Description Эндпоинт для подвтерждения почты по ссылке. Через query параметры получаются
// @Description code и username, а затем подтверждается эта почта.
// @Accept json
// @Produce json
// @Param input query request.EmailConfirmation true "Confirm"
// @Success 200 {string} ok
// @Failure 400 {string} bad request
// @Failure 500 {string} internal server error
// @Router /api/auth/confirm [get]
func (h *AuthHandler) ConfirmEmail(c *gin.Context) {
	code := c.Query("code")
	username := c.Query("username")

	confirmationDto := dto.EmailConfirmation{
		Code:     code,
		Username: username,
	}

	if err := confirmationDto.Validate(); err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	err := h.service.ConfirmAccountEmail(c.Request.Context(), confirmationDto.Code, confirmationDto.Username)

	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
