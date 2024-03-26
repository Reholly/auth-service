package handler

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/repository"
	"auth-service/internal/server"
	"auth-service/internal/server/request/dto"
	"auth-service/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AdminHandler struct {
	repository *repository.RepositoryManager
	service    *service.ServiceManager
}

func NewAdminHandler(
	repositoryManager *repository.RepositoryManager,
	serviceManager *service.ServiceManager) *AdminHandler {
	return &AdminHandler{
		repository: repositoryManager,
		service:    serviceManager,
	}
}

// CreateModerator godoc
// @Summary Endpoint for moderator creating.
// @Description Этот эндпоинт нужен для накидывания Claim модератора на пользователя.
// @Description С эндпоинта можно получить 200, 400, 500 статус-коды.
// @Description Ошибки с эндпоинта: Bad Credentials, Internal Server Error
// @Accept json
// @Produce json
// @Param input body request.Username true "Username"
// @Success 200 {string} ok
// @Failure 400 {string} bad request: error in credentials
// @Failure 500 {string} server error: reset error
// @Router /api/auth/admin/createmoder [put]
func (h *AdminHandler) CreateModerator(c *gin.Context) {
	claims, err := h.authorize(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}

	if !h.service.TokenService.IsAdmin(claims) {
		c.AbortWithStatusJSON(http.StatusForbidden, "not admin")
		return
	}

	var username dto.Username
	if err := c.ShouldBindQuery(&username); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, "no username")
		return
	}

	err = h.service.AdminService.CreateModerator(c.Request.Context(), username.Username)
	if err != nil {
		apiError := server.NewAPIError(err)
		c.AbortWithStatusJSON(apiError.Code, apiError.Message)
		return
	}

	c.JSON(http.StatusOK, claims)
}

func (h *AdminHandler) DeleteModerator(c *gin.Context) {

}

func (h *AdminHandler) BanUser(c *gin.Context) {

}

func (h *AdminHandler) UnbanUser(c *gin.Context) {

}

func (h *AdminHandler) authorize(c *gin.Context) ([]entity.Claim, error) {
	bearerHeader := c.Request.Header.Get("Authorization")
	if bearerHeader == "" {
		return nil, errors.New("empty header")
	}
	token := strings.Split(bearerHeader, " ")[1]
	return h.service.TokenService.ParseClaims(token)
}
