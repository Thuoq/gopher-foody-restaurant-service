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

type UpdateRestaurantInput struct {
	PublicID    string
	OwnerID     string
	Name        string
	Address     string
	Description string
	LogoURL     string
	BannerURL   string
}

type ListRestaurantsInput struct {
	OwnerID string
	Search  string
	Status  string
	Page    int
	Limit   int
}

type IRestaurantUseCase interface {
	CreateRestaurant(ctx context.Context, input CreateRestaurantInput) (*domain.Restaurant, error)
	ListRestaurants(ctx context.Context, input ListRestaurantsInput) ([]domain.Restaurant, int64, error)
	GetRestaurantDetail(ctx context.Context, publicID string) (*domain.Restaurant, error)
	UpdateRestaurant(ctx context.Context, input UpdateRestaurantInput) (*domain.Restaurant, error)
	DeleteRestaurant(ctx context.Context, ownerID string, publicID string) error
}
