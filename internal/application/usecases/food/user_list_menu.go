package food

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type userListMenuUseCase struct {
	foodRepo       ports.FoodRepository
	restaurantRepo ports.RestaurantRepository
}

func NewUserListMenuUseCase(foodRepo ports.FoodRepository, restaurantRepo ports.RestaurantRepository) ports.IUserListMenuUseCase {
	return &userListMenuUseCase{
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *userListMenuUseCase) Execute(ctx context.Context, restaurantPublicID string) ([]domain.Food, error) {
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, restaurantPublicID)
	if err != nil {
		return nil, err
	}
	return uc.foodRepo.ListByRestaurant(ctx, restaurant.ID)
}
