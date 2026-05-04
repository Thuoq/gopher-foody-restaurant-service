package usecases

import (
	"context"
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"github.com/google/uuid"
)

type foodUseCase struct {
	foodRepo       ports.FoodRepository
	restaurantRepo ports.RestaurantRepository
}

func NewFoodUseCase(foodRepo ports.FoodRepository, restaurantRepo ports.RestaurantRepository) ports.IFoodUseCase {
	return &foodUseCase{
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *foodUseCase) CreateFood(ctx context.Context, ownerID string, input ports.CreateFoodInput) (*domain.Food, error) {
	// 1. Verify restaurant existence and ownership
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
		img := &domain.FoodImage{
			FoodID:   food.ID,
			ImageURL: imgURL,
		}
		_ = uc.foodRepo.AddImage(ctx, img)
	}

	return food, nil
}

func (uc *foodUseCase) ListRestaurantMenu(ctx context.Context, restaurantPublicID string) ([]domain.Food, error) {
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, restaurantPublicID)
	if err != nil {
		return nil, err
	}
	return uc.foodRepo.ListByRestaurant(ctx, restaurant.ID)
}

func (uc *foodUseCase) UpdateFood(ctx context.Context, ownerID string, input ports.UpdateFoodInput) (*domain.Food, error) {
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

	// 3. Strict Validation
	if input.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}

	// 4. Update fields
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

func (uc *foodUseCase) DeleteFood(ctx context.Context, ownerID string, publicID string) error {
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
