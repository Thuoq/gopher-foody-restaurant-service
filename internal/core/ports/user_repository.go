package ports

import (
	"context"

	"gopher-identity-service/internal/core/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	// Add other methods like Create, GetByEmail etc. as needed
}
