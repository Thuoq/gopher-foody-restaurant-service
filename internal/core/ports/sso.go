package ports

import (
	"context"

	"gopher-identity-service/internal/core/domain"
)

type SSOUseCase interface {
	GetUserProfile(ctx context.Context, userID int64) (*domain.User, error)
}
