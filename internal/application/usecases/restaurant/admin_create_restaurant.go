package restaurant

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"github.com/google/uuid"
)

type adminCreateRestaurantUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewAdminCreateRestaurantUseCase(restaurantRepo ports.RestaurantRepository) ports.IAdminCreateRestaurantUseCase {
	return &adminCreateRestaurantUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminCreateRestaurantUseCase) Execute(ctx context.Context, input ports.CreateRestaurantInput) (*domain.Restaurant, error) {
	restaurant := &domain.Restaurant{
		PublicID:    uuid.New().String(),
		OwnerID:     input.OwnerID,
		Name:        input.Name,
		Address:     input.Address,
		Description: input.Description,
		LogoURL:     input.LogoURL,
		BannerURL:   input.BannerURL,
		Status:      "active",
	}

	if err := uc.restaurantRepo.Create(ctx, restaurant); err != nil {
		return nil, err
	}

	return restaurant, nil
}
