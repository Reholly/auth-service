package application

import (
	"auth-service/internal/config"
	"auth-service/internal/service"
	"auth-service/internal/storage/repositories"
	"auth-service/lib/db"
	"context"
)

type Application struct {
	config config.Config
}

func (a *Application) Run() {
	ctx := context.Background()

	postgresAdapter := db.NewPostgresAdapter()

	_, err := postgresAdapter.Connect(ctx, a.config.ConnectionString)
	if err != nil {
		panic(err)
	}

	accountRepository := repositories.NewAccountRepository(*postgresAdapter)
	claimRepository := repositories.NewClaimRepository(*postgresAdapter)

	repositoryManager := repositories.NewRepositoryManager(accountRepository, claimRepository)

	adminService := service.NewAdminService(repositoryManager)
	mailService := service.NewMailService(a.config)
	tokenService := service.NewTokenService(a.config)
	authService := service.NewAuthService(a.config, repositoryManager, mailService, tokenService)

	serviceManager := service.NewServiceManager(mailService, adminService, authService, tokenService)
}
