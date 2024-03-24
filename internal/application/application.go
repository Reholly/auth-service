package application

import (
	"auth-service/config"
	"auth-service/internal/repository"
	"auth-service/internal/repository/inMemory"
	repositories2 "auth-service/internal/repository/postgres/implementation"
	"auth-service/internal/server/handler"
	"auth-service/internal/server/router"
	"auth-service/internal/service"
	"auth-service/internal/service/implementation"
	"auth-service/lib/db"
	"context"
	"fmt"
	"time"
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
	codeRepository := inMemory.NewCache(
		time.Duration(a.config.CodeExpirationTime),
		time.Duration(a.config.CodeExpirationTime))

	repositoryManager := repository.NewRepositoryManager(accountRepository, claimRepository, codeRepository)

	mailService := implementation.NewMailService(a.config)
	adminService := implementation.NewAdminService(repositoryManager, mailService)
	tokenService := implementation.NewTokenService(a.config)
	authService := implementation.NewAuthService(a.config, repositoryManager, mailService, tokenService)

	serviceManager := service.NewServiceManager(mailService, adminService, authService, tokenService)
	fmt.Println(serviceManager)

	router := router.NewRouter(a.config)

	authHandler := handler.NewAuthHandler(serviceManager, repositoryManager)
	accountHandler := handler.NewAccountHandler(serviceManager, repositoryManager)

	err = router.Run(authHandler, accountHandler)

	if err != nil {
		panic(err)
	}
}
