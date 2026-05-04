package food_category

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
)

type userListFoodCategoriesUseCase struct {
	categoryRepo ports.FoodCategoryRepository
}

func NewUserListFoodCategoriesUseCase(categoryRepo ports.FoodCategoryRepository) ports.IUserListFoodCategoriesUseCase {
	return &userListFoodCategoriesUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *userListFoodCategoriesUseCase) Execute(ctx context.Context) ([]domain.FoodCategory, error) {
	return uc.categoryRepo.List(ctx)
}
