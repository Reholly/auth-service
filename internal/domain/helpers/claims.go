package domain

import "auth-service/internal/domain/entity"

var (
	AdminRole = domain.Claim{
		Title: "role",
		Value: "admin",
	}
	StudentRole = domain.Claim{
		Title: "role",
		Value: "student",
	}
	ModeratorRole = domain.Claim{
		Title: "role",
		Value: "moderator",
	}
)
