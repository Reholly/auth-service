package service

import "auth-service/internal/domain"

type ServiceManager struct {
	domain.MailService
	domain.AdminService
	domain.AuthService
	domain.TokenService
}

func NewServiceManager(
	mailService domain.MailService,
	adminService domain.AdminService,
	authService domain.AuthService,
	tokenService domain.TokenService) *ServiceManager {
	return &ServiceManager{
		MailService:  mailService,
		AdminService: adminService,
		AuthService:  authService,
		TokenService: tokenService,
	}
}
