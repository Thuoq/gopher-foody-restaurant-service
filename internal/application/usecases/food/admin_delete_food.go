package food

import (
	"context"
	"errors"
	"gopher-restaurant-service/internal/core/ports"
)

type adminDeleteFoodUseCase struct {
	foodRepo       ports.FoodRepository
	restaurantRepo ports.RestaurantRepository
}

func NewAdminDeleteFoodUseCase(foodRepo ports.FoodRepository, restaurantRepo ports.RestaurantRepository) ports.IAdminDeleteFoodUseCase {
	return &adminDeleteFoodUseCase{
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminDeleteFoodUseCase) Execute(ctx context.Context, ownerID string, publicID string) error {
	// 1. Fetch existing food
	food, err := uc.foodRepo.GetByPublicID(ctx, publicID)
	if err != nil {
		return err
	}

	// 2. Fetch restaurant to verify ownership
	restaurant, err := uc.restaurantRepo.GetByID(ctx, food.RestaurantID)
	if err != nil {
		return err
	}
	if restaurant.OwnerID != ownerID {
		return errors.New("unauthorized: you do not own the restaurant this food belongs to")
	}

	return uc.foodRepo.Delete(ctx, publicID)
}
