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
	Images             []string
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

type IAdminCreateFoodUseCase interface {
	Execute(ctx context.Context, ownerID string, input CreateFoodInput) (*domain.Food, error)
}

type IAdminUpdateFoodUseCase interface {
	Execute(ctx context.Context, ownerID string, input UpdateFoodInput) (*domain.Food, error)
}

type IAdminDeleteFoodUseCase interface {
	Execute(ctx context.Context, ownerID string, publicID string) error
}

type IUserListMenuUseCase interface {
	Execute(ctx context.Context, restaurantPublicID string) ([]domain.Food, error)
}
