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

type IAdminCreateRestaurantUseCase interface {
	Execute(ctx context.Context, input CreateRestaurantInput) (*domain.Restaurant, error)
}

type IAdminUpdateRestaurantUseCase interface {
	Execute(ctx context.Context, input UpdateRestaurantInput) (*domain.Restaurant, error)
}

type IAdminDeleteRestaurantUseCase interface {
	Execute(ctx context.Context, ownerID string, publicID string) error
}

type IAdminListMyRestaurantsUseCase interface {
	Execute(ctx context.Context, input ListRestaurantsInput) ([]domain.Restaurant, int64, error)
}

type IUserListRestaurantsUseCase interface {
	Execute(ctx context.Context, input ListRestaurantsInput) ([]domain.Restaurant, int64, error)
}

type IUserGetRestaurantDetailUseCase interface {
	Execute(ctx context.Context, publicID string) (*domain.Restaurant, error)
}
