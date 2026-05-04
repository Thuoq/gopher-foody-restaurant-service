package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type CreateCategoryInput struct {
	Name    string
	IconURL string
}

type UpdateCategoryInput struct {
	ID      uint
	Name    string
	IconURL string
}

type IAdminCreateCategoryUseCase interface {
	Execute(ctx context.Context, input CreateCategoryInput) (*domain.Category, error)
}

type IAdminUpdateCategoryUseCase interface {
	Execute(ctx context.Context, input UpdateCategoryInput) (*domain.Category, error)
}

type IAdminDeleteCategoryUseCase interface {
	Execute(ctx context.Context, id uint) error
}

type IUserListCategoriesUseCase interface {
	Execute(ctx context.Context) ([]domain.Category, error)
}
