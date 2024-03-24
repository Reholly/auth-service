package helpers

import "auth-service/internal/domain/entity"

var (
	AdminRole = entity.Claim{
		Title: "role",
		Value: "admin",
	}
	StudentRole = entity.Claim{
		Title: "role",
		Value: "student",
	}
	ModeratorRole = entity.Claim{
		Title: "role",
		Value: "moderator",
	}
)
