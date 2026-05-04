package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type FoodCategoryRepository interface {
	Create(ctx context.Context, category *domain.FoodCategory) error
	GetByID(ctx context.Context, id uint) (*domain.FoodCategory, error)
	List(ctx context.Context) ([]domain.FoodCategory, error)
	Update(ctx context.Context, category *domain.FoodCategory) error
	Delete(ctx context.Context, id uint) error
}
