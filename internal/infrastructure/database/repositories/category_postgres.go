package repositories

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"gorm.io/gorm"
)

type categoryPostgresRepository struct {
	db *gorm.DB
}

func NewCategoryPostgresRepository(db *gorm.DB) ports.CategoryRepository {
	return &categoryPostgresRepository{
		db: db,
	}
}

func (r *categoryPostgresRepository) Create(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryPostgresRepository) GetByID(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	return &category, err
}

func (r *categoryPostgresRepository) List(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.WithContext(ctx).Order("name ASC").Find(&categories).Error
	return categories, err
}

func (r *categoryPostgresRepository) Update(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *categoryPostgresRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Category{}, id).Error
}
