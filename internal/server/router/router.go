package router

import (
	"auth-service/internal/config"
	"auth-service/internal/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	address string
}

func NewRouter(config config.Config) *Router {
	return &Router{
		address: config.Address,
	}
}

func (r *Router) Run(authHandler handler.AuthHandler) {
	/*	go func() {
			_ = r.gin.Run(r.address)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit*/

	g := gin.New()

	group := g.Group("/api")

	auth := group.Group("/auth")
	{
		auth.GET("/login", authHandler.LogIn)
		auth.POST("/register", authHandler.Register)
		auth.GET("/reset", authHandler.ResetPassword)
		auth.GET("/confirm", authHandler.ConfirmEmail)
	}

	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

}
