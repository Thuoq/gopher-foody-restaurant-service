package usecases

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"github.com/google/uuid"
)

type restaurantUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewRestaurantUseCase(restaurantRepo ports.RestaurantRepository) ports.IRestaurantUseCase {
	return &restaurantUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *restaurantUseCase) CreateRestaurant(ctx context.Context, input ports.CreateRestaurantInput) (*domain.Restaurant, error) {
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

func (uc *restaurantUseCase) GetMyRestaurants(ctx context.Context, ownerID string) ([]domain.Restaurant, error) {
	return uc.restaurantRepo.ListByOwner(ctx, ownerID)
}

func (uc *restaurantUseCase) ListRestaurants(ctx context.Context) ([]domain.Restaurant, error) {
	return uc.restaurantRepo.ListAll(ctx)
}

func (uc *restaurantUseCase) GetRestaurantDetail(ctx context.Context, publicID string) (*domain.Restaurant, error) {
	return uc.restaurantRepo.GetByPublicID(ctx, publicID)
}
