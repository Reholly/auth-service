package handler

import (
	"auth-service/internal/repository"
	"auth-service/internal/server"
	"auth-service/internal/server/request/dto"
	"auth-service/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	service    *service.ServiceManager
	repository *repository.RepositoryManager
}

func NewAccountHandler(service *service.ServiceManager, repository *repository.RepositoryManager) *AccountHandler {
	return &AccountHandler{
		service:    service,
		repository: repository,
	}
}

// ConfirmResetPassword godoc
// @Summary Endpoint for password reset.
// @Description Этот эндпоинт нужен для сброса пароля. Пользователь вводит одноразовый код, высланный ранее на почту.
// @Description В теле запроса пользователь должен ввести код, пришедший на почту, новый пароль и повторенный новый пароль.
// @Description С эндпоинта можно получить 200, 400, 500 статус-коды.
// @Description Ошибки с эндпоинта: Bad Credentials, Internal Server Error
// @Accept json
// @Produce json
// @Param input body request.PasswordResetConfirmation true "Reset"
// @Success 200 {string} ok
// @Failure 400 {string} bad request: error in credentials
// @Failure 500 {string} server error: reset error
// @Router /api/auth/account/confirmreset [put]
func (h *AccountHandler) ConfirmResetPassword(c *gin.Context) {

	var passwordReset dto.PasswordResetConfirmation
	if err := c.ShouldBindJSON(&passwordReset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := passwordReset.Validate(); err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	err := h.service.ResetPassword(c.Request.Context(), passwordReset.Code, passwordReset.Password)

	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// SendResetPasswordCode godoc
// @Summary Endpoint for password reset request.
// @Description Этот эндпоинт нужен для запроса сброса пароля. Пользователь вводит Email, по этой почте ищется пользователь и генерируется одноразовый код.
// @Description С эндпоинта можно получить 200, 400, 500 статус-коды.
// @Description Ошибки с эндпоинта: Bad Credentials, Internal Server Error
// @Accept json
// @Produce json
// @Param input body request.PasswordResetRequest true "PasswordResetConfirmation"
// @Success 200 {string} ok
// @Failure 400 {string} bad request: error in credentials
// @Failure 500 {string} server error: reset error
// @Router /api/auth/account/sendresetcode [put]
func (h *AccountHandler) SendResetPasswordCode(c *gin.Context) {
	var passwordReset dto.PasswordResetRequest
	if err := c.ShouldBindJSON(&passwordReset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := passwordReset.Validate(); err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	account, err := h.repository.AccountRepository.GetAccountByEmail(c.Request.Context(), passwordReset.Email)
	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}
	code, _ := h.service.AuthService.GenerateResetPasswordCode(c.Request.Context(), account.Username)

	err = h.service.SendMail(c.Request.Context(),
		account.Email,
		"Запрос на сброс пароля",
		fmt.Sprintf("Для сброса пароля вам выслан код: %s. Код действителен в течение 20 минут."+
			" Если это были не вы, обратитесь к администрации KForge.", code),
	)

	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
