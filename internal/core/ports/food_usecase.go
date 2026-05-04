package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type CreateFoodInput struct {
	RestaurantPublicID string
	CategoryID         uint
	Name               string
	Description        string
	Price              float64
	Quantity           int
	Images             []string // List of image URLs
}

type UpdateFoodInput struct {
	PublicID    string
	CategoryID  uint
	Name        string
	Description string
	Price       float64
	Quantity    int
	Status      string
	Images      []string
}

type IFoodUseCase interface {
	CreateFood(ctx context.Context, ownerID string, input CreateFoodInput) (*domain.Food, error)
	ListRestaurantMenu(ctx context.Context, restaurantPublicID string) ([]domain.Food, error)
	UpdateFood(ctx context.Context, ownerID string, input UpdateFoodInput) (*domain.Food, error)
	DeleteFood(ctx context.Context, ownerID string, publicID string) error
}
