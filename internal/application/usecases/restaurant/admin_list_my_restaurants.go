package restaurant

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/pkg/utils"
)

type adminListMyRestaurantsUseCase struct {
	restaurantRepo ports.RestaurantRepository
}

func NewAdminListMyRestaurantsUseCase(restaurantRepo ports.RestaurantRepository) ports.IAdminListMyRestaurantsUseCase {
	return &adminListMyRestaurantsUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *adminListMyRestaurantsUseCase) Execute(ctx context.Context, input ports.ListRestaurantsInput) ([]domain.Restaurant, int64, error) {
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
