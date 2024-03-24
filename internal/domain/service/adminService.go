package domain

import "context"

type AdminService interface {
	BanUser(ctx context.Context, username, reason string) error
	UnbanUser(ctx context.Context, username string) error
	CreateModerator(ctx context.Context, username string) error
	DeleteModerator(ctx context.Context, username string) error
}
