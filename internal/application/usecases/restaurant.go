package usecases

import (
	"context"
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/pkg/utils"

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

func (uc *restaurantUseCase) ListRestaurants(ctx context.Context, input ports.ListRestaurantsInput) ([]domain.Restaurant, int64, error) {
	p := utils.Pagination{
		Page:  input.Page,
		Limit: input.Limit,
	}

	filter := ports.ListRestaurantsFilter{
		OwnerID: input.OwnerID,
		Search:  input.Search,
		Status:  input.Status,
		Offset:  p.GetOffset(),
		Limit:   p.GetLimit(),
	}

	return uc.restaurantRepo.List(ctx, filter)
}

func (uc *restaurantUseCase) GetRestaurantDetail(ctx context.Context, publicID string) (*domain.Restaurant, error) {
	return uc.restaurantRepo.GetByPublicID(ctx, publicID)
}

func (uc *restaurantUseCase) UpdateRestaurant(ctx context.Context, input ports.UpdateRestaurantInput) (*domain.Restaurant, error) {
	// 1. Fetch existing restaurant
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, input.PublicID)
	if err != nil {
		return nil, err
	}

	// 2. Authorization check
	if restaurant.OwnerID != input.OwnerID {
		return nil, errors.New("unauthorized: you do not own this restaurant")
	}

	// 3. Strict Validation (Example: Name cannot be empty)
	if input.Name == "" {
		return nil, errors.New("restaurant name is required")
	}

	// 4. Update fields
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

func (uc *restaurantUseCase) DeleteRestaurant(ctx context.Context, ownerID string, publicID string) error {
	// 1. Fetch existing restaurant
	restaurant, err := uc.restaurantRepo.GetByPublicID(ctx, publicID)
	if err != nil {
		return err
	}

	// 2. Authorization check
	if restaurant.OwnerID != ownerID {
		return errors.New("unauthorized: you do not own this restaurant")
	}

	return uc.restaurantRepo.Delete(ctx, publicID)
}
