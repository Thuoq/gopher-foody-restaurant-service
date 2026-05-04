package food_category

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type adminCreateFoodCategoryUseCase struct {
	categoryRepo ports.FoodCategoryRepository
}

func NewAdminCreateFoodCategoryUseCase(categoryRepo ports.FoodCategoryRepository) ports.IAdminCreateFoodCategoryUseCase {
	return &adminCreateFoodCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *adminCreateFoodCategoryUseCase) Execute(ctx context.Context, input ports.CreateFoodCategoryInput) (*domain.FoodCategory, error) {
	category := &domain.FoodCategory{
		Name:    input.Name,
		IconURL: input.IconURL,
	}

	if err := uc.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
