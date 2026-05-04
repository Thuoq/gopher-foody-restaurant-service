package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type ListRestaurantsFilter struct {
	OwnerID string
	Search  string
	Status  string
	Offset  int
	Limit   int
}

type RestaurantRepository interface {
	Create(ctx context.Context, restaurant *domain.Restaurant) error
	GetByID(ctx context.Context, id uint) (*domain.Restaurant, error)
	GetByPublicID(ctx context.Context, publicID string) (*domain.Restaurant, error)
	List(ctx context.Context, filter ListRestaurantsFilter) ([]domain.Restaurant, int64, error)
	Update(ctx context.Context, restaurant *domain.Restaurant) error
	Delete(ctx context.Context, publicID string) error

	// Image related
	AddImage(ctx context.Context, image *domain.RestaurantImage) error
}
