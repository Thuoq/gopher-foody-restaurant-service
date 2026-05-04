package ports

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
)

type FoodRepository interface {
	Create(ctx context.Context, food *domain.Food) error
	GetByPublicID(ctx context.Context, publicID string) (*domain.Food, error)
	ListByRestaurant(ctx context.Context, restaurantID uint) ([]domain.Food, error)
	Update(ctx context.Context, food *domain.Food) error
	
	// Image related
	AddImage(ctx context.Context, image *domain.FoodImage) error
}

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	List(ctx context.Context) ([]domain.Category, error)
	GetByID(ctx context.Context, id uint) (*domain.Category, error)
}
