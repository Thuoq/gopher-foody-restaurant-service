package food

import (
	"context"
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"github.com/google/uuid"
)

type adminCreateFoodUseCase struct {
	foodRepo       ports.FoodRepository
	restaurantRepo ports.RestaurantRepository
}

func NewAdminCreateFoodUseCase(foodRepo ports.FoodRepository, restaurantRepo ports.RestaurantRepository) ports.IAdminCreateFoodUseCase {
	return &adminCreateFoodUseCase{
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminCreateFoodUseCase) Execute(ctx context.Context, ownerID string, input ports.CreateFoodInput) (*domain.Food, error) {
	// 1. Fetch restaurant and verify ownership
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, input.RestaurantPublicID)
	if err != nil {
		return nil, err
	}

	if restaurant.OwnerID != ownerID {
		return nil, errors.New("unauthorized: you do not own this restaurant")
	}

	// 2. Create food
	food := &domain.Food{
		PublicID:     uuid.New().String(),
		RestaurantID: restaurant.ID,
		CategoryID:   input.CategoryID,
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		Quantity:     input.Quantity,
		Status:       "available",
	}

	if err := uc.foodRepo.Create(ctx, food); err != nil {
		return nil, err
	}

	// 3. Add images
	for _, imgURL := range input.Images {
		foodImage := &domain.FoodImage{
			FoodID: food.ID,
			URL:    imgURL,
		}
		_ = uc.foodRepo.AddImage(ctx, foodImage)
	}

	return food, nil
}
