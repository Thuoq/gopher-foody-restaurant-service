package food_category

import (
	"context"
	"gopher-restaurant-service/internal/core/ports"
)

type adminDeleteFoodCategoryUseCase struct {
	categoryRepo ports.FoodCategoryRepository
}

func NewAdminDeleteFoodCategoryUseCase(categoryRepo ports.FoodCategoryRepository) ports.IAdminDeleteFoodCategoryUseCase {
	return &adminDeleteFoodCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *adminDeleteFoodCategoryUseCase) Execute(ctx context.Context, id uint) error {
	return uc.categoryRepo.Delete(ctx, id)
}
