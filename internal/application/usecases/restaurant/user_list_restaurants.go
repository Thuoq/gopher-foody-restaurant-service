package restaurant

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/pkg/utils"
)

type userListRestaurantsUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewUserListRestaurantsUseCase(restaurantRepo ports.RestaurantRepository) ports.IUserListRestaurantsUseCase {
	return &userListRestaurantsUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *userListRestaurantsUseCase) Execute(ctx context.Context, input ports.ListRestaurantsInput) ([]domain.Restaurant, int64, error) {
	p := utils.Pagination{
		Page:  input.Page,
		Limit: input.Limit,
	}

	filter := ports.ListRestaurantsFilter{
		Search: input.Search,
		Status: "active", // Customers only see active restaurants
		Offset: p.GetOffset(),
		Limit:  p.GetLimit(),
	}

	return uc.restaurantRepo.List(ctx, filter)
}
