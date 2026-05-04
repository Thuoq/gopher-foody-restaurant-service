package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type CreateRestaurantInput struct {
	OwnerID     string
	Name        string
	Address     string
	Description string
	LogoURL     string
	BannerURL   string
}

type IRestaurantUseCase interface {
	CreateRestaurant(ctx context.Context, input CreateRestaurantInput) (*domain.Restaurant, error)
	GetMyRestaurants(ctx context.Context, ownerID string) ([]domain.Restaurant, error)
	ListRestaurants(ctx context.Context) ([]domain.Restaurant, error)
	GetRestaurantDetail(ctx context.Context, publicID string) (*domain.Restaurant, error)
}
