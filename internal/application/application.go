package application

import (
	"auth-service/internal/config"
	"auth-service/internal/server/handler"
	"auth-service/internal/server/router"
	"auth-service/internal/service"
	repositories2 "auth-service/internal/storage/postgres/repositories"
	"auth-service/lib/db"
	"context"
	"fmt"
)

type Application struct {
	config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{config: config}
}

func (a *Application) Run() {
	ctx := context.Background()

	postgresAdapter := db.NewPostgresAdapter()

	_, err := postgresAdapter.Connect(ctx, a.config.ConnectionString)
	if err != nil {
		panic(err)
	}

	accountRepository := repositories2.NewAccountRepository(postgresAdapter)
	claimRepository := repositories2.NewClaimRepository(postgresAdapter)

	repositoryManager := repositories2.NewRepositoryManager(accountRepository, claimRepository)

	mailService := service.NewMailService(a.config)
	adminService := service.NewAdminService(repositoryManager, mailService)
	tokenService := service.NewTokenService(a.config)
	authService := service.NewAuthService(a.config, repositoryManager, mailService, tokenService)

	serviceManager := service.NewServiceManager(mailService, adminService, authService, tokenService)
	fmt.Println(serviceManager)
	router := router.NewRouter(a.config)
	authHandler := handler.NewAuthHandler(serviceManager, repositoryManager)
	router.Run(authHandler)
}
