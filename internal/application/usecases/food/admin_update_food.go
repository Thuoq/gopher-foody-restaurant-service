package food

import (
	"context"
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type adminUpdateFoodUseCase struct {
	foodRepo       ports.FoodRepository
	restaurantRepo ports.RestaurantRepository
}

func NewAdminUpdateFoodUseCase(foodRepo ports.FoodRepository, restaurantRepo ports.RestaurantRepository) ports.IAdminUpdateFoodUseCase {
	return &adminUpdateFoodUseCase{
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminUpdateFoodUseCase) Execute(ctx context.Context, ownerID string, input ports.UpdateFoodInput) (*domain.Food, error) {
	// 1. Fetch existing food
	food, err := uc.foodRepo.GetByPublicID(ctx, input.PublicID)
	if err != nil {
		return nil, err
	}

	// 2. Fetch restaurant to verify ownership
	restaurant, err := uc.restaurantRepo.GetByID(ctx, food.RestaurantID)
	if err != nil {
		return nil, err
	}
	if restaurant.OwnerID != ownerID {
		return nil, errors.New("unauthorized: you do not own the restaurant this food belongs to")
	}

	// 3. Update fields
	food.Name = input.Name
	food.Description = input.Description
	food.Price = input.Price
	food.Quantity = input.Quantity
	food.CategoryID = input.CategoryID
	food.Status = input.Status

	if err := uc.foodRepo.Update(ctx, food); err != nil {
		return nil, err
	}

	return food, nil
}
