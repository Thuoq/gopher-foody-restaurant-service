package repositories

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"gorm.io/gorm"
)

type foodCategoryPostgresRepository struct {
	db *gorm.DB
}

func NewFoodCategoryPostgresRepository(db *gorm.DB) ports.FoodCategoryRepository {
	return &foodCategoryPostgresRepository{
		db: db,
	}
}

func (r *foodCategoryPostgresRepository) Create(ctx context.Context, category *domain.FoodCategory) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *foodCategoryPostgresRepository) GetByID(ctx context.Context, id uint) (*domain.FoodCategory, error) {
	var category domain.FoodCategory
	err := r.db.WithContext(ctx).First(&category, id).Error
	return &category, err
}

func (r *foodCategoryPostgresRepository) List(ctx context.Context) ([]domain.FoodCategory, error) {
	var categories []domain.FoodCategory
	err := r.db.WithContext(ctx).Order("name ASC").Find(&categories).Error
	return categories, err
}

func (r *foodCategoryPostgresRepository) Update(ctx context.Context, category *domain.FoodCategory) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *foodCategoryPostgresRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.FoodCategory{}, id).Error
}
