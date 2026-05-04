package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type CreateFoodCategoryInput struct {
	Name    string
	IconURL string
}

type UpdateFoodCategoryInput struct {
	ID      uint
	Name    string
	IconURL string
}

type IAdminCreateFoodCategoryUseCase interface {
	Execute(ctx context.Context, input CreateFoodCategoryInput) (*domain.FoodCategory, error)
}

type IAdminUpdateFoodCategoryUseCase interface {
	Execute(ctx context.Context, input UpdateFoodCategoryInput) (*domain.FoodCategory, error)
}

type IAdminDeleteFoodCategoryUseCase interface {
	Execute(ctx context.Context, id uint) error
}

type IUserListFoodCategoriesUseCase interface {
	Execute(ctx context.Context) ([]domain.FoodCategory, error)
}
