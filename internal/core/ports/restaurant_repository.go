package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type RestaurantRepository interface {
	Create(ctx context.Context, restaurant *domain.Restaurant) error
	GetByPublicID(ctx context.Context, publicID string) (*domain.Restaurant, error)
	ListByOwner(ctx context.Context, ownerID string) ([]domain.Restaurant, error)
	ListAll(ctx context.Context) ([]domain.Restaurant, error)
	Update(ctx context.Context, restaurant *domain.Restaurant) error

	// Image related
	AddImage(ctx context.Context, image *domain.RestaurantImage) error
}
