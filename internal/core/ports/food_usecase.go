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

type IFoodUseCase interface {
	CreateFood(ctx context.Context, ownerID string, input CreateFoodInput) (*domain.Food, error)
	ListRestaurantMenu(ctx context.Context, restaurantPublicID string) ([]domain.Food, error)
}
