package restaurant

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type userGetRestaurantDetailUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewUserGetRestaurantDetailUseCase(restaurantRepo ports.RestaurantRepository) ports.IUserGetRestaurantDetailUseCase {
	return &userGetRestaurantDetailUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *userGetRestaurantDetailUseCase) Execute(ctx context.Context, publicID string) (*domain.Restaurant, error) {
	return uc.restaurantRepo.GetByPublicID(ctx, publicID)
}
