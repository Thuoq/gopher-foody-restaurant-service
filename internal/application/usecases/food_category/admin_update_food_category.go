package food_category

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type adminUpdateFoodCategoryUseCase struct {
	categoryRepo ports.FoodCategoryRepository
}

func NewAdminUpdateFoodCategoryUseCase(categoryRepo ports.FoodCategoryRepository) ports.IAdminUpdateFoodCategoryUseCase {
	return &adminUpdateFoodCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *adminUpdateFoodCategoryUseCase) Execute(ctx context.Context, input ports.UpdateFoodCategoryInput) (*domain.FoodCategory, error) {
	category, err := uc.categoryRepo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	category.Name = input.Name
	category.IconURL = input.IconURL

	if err := uc.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
