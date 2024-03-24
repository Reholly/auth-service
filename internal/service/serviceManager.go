package service

import (
	"auth-service/internal/domain/service"
)

type ServiceManager struct {
	service.MailService
	service.AdminService
	service.AuthService
	service.TokenService
}

func NewServiceManager(
	mailService service.MailService,
	adminService service.AdminService,
	authService service.AuthService,
	tokenService service.TokenService) *ServiceManager {
	return &ServiceManager{
		MailService:  mailService,
		AdminService: adminService,
		AuthService:  authService,
		TokenService: tokenService,
	}
}
