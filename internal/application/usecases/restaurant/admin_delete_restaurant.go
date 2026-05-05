package restaurant

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type adminDeleteRestaurantUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewAdminDeleteRestaurantUseCase(restaurantRepo ports.RestaurantRepository) ports.IAdminDeleteRestaurantUseCase {
	return &adminDeleteRestaurantUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminDeleteRestaurantUseCase) Execute(ctx context.Context, ownerID string, publicID string) error {
	// 1. Fetch existing restaurant
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, publicID)
	if err != nil {
		return err
	}

	// 2. Authorization check
	if restaurant.OwnerID != ownerID {
		return domain.ErrUnauthorized
	}

	return uc.restaurantRepo.Delete(ctx, publicID)
}
