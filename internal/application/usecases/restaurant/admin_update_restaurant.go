package restaurant

import (
	"context"
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type adminUpdateRestaurantUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewAdminUpdateRestaurantUseCase(restaurantRepo ports.RestaurantRepository) ports.IAdminUpdateRestaurantUseCase {
	return &adminUpdateRestaurantUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminUpdateRestaurantUseCase) Execute(ctx context.Context, input ports.UpdateRestaurantInput) (*domain.Restaurant, error) {
	// 1. Fetch existing restaurant
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, input.PublicID)
	if err != nil {
		return nil, err
	}

	// 2. Authorization check
	if restaurant.OwnerID != input.OwnerID {
		return nil, errors.New("unauthorized: you do not own this restaurant")
	}

	// 3. Update fields
	restaurant.Name = input.Name
	restaurant.Address = input.Address
	restaurant.Description = input.Description
	restaurant.LogoURL = input.LogoURL
	restaurant.BannerURL = input.BannerURL

	if err := uc.restaurantRepo.Update(ctx, restaurant); err != nil {
		return nil, err
	}

	return restaurant, nil
}
