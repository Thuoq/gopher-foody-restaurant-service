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
		return nil, domain.ErrUnauthorized
	}

	// 3. Update fields (Partial update)
	if input.Name != nil {
		restaurant.Name = *input.Name
	}
	if input.Address != nil {
		restaurant.Address = *input.Address
	}
	if input.Description != nil {
		restaurant.Description = *input.Description
	}
	if input.LogoURL != nil {
		restaurant.LogoURL = *input.LogoURL
	}
	if input.BannerURL != nil {
		restaurant.BannerURL = *input.BannerURL
	}

	if err := uc.restaurantRepo.Update(ctx, restaurant); err != nil {
		return nil, err
	}

	return restaurant, nil
}
