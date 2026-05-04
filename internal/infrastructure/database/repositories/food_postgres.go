package repositories

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"gorm.io/gorm"
)

type foodPostgresRepository struct {
	db *gorm.DB
}

func NewFoodPostgresRepository(db *gorm.DB) ports.FoodRepository {
	return &foodPostgresRepository{
		db: db,
	}
}

func (r *foodPostgresRepository) Create(ctx context.Context, food *domain.Food) error {
	return r.db.WithContext(ctx).Create(food).Error
}

func (r *foodPostgresRepository) GetByPublicID(ctx context.Context, publicID string) (*domain.Food, error) {
	var food domain.Food
	err := r.db.WithContext(ctx).Preload("Images").Where("public_id = ?", publicID).First(&food).Error
	return &food, err
}

func (r *foodPostgresRepository) ListByRestaurant(ctx context.Context, restaurantID uint) ([]domain.Food, error) {
	var foods []domain.Food
	err := r.db.WithContext(ctx).Preload("Images").Where("restaurant_id = ?", restaurantID).Find(&foods).Error
	return foods, err
}

func (r *foodPostgresRepository) Update(ctx context.Context, food *domain.Food) error {
	return r.db.WithContext(ctx).Save(food).Error
}

func (r *foodPostgresRepository) Delete(ctx context.Context, publicID string) error {
	return r.db.WithContext(ctx).Where("public_id = ?", publicID).Delete(&domain.Food{}).Error
}

func (r *foodPostgresRepository) AddImage(ctx context.Context, image *domain.FoodImage) error {
	return r.db.WithContext(ctx).Create(image).Error
}

// Category Repository Implementation
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

func (r *categoryPostgresRepository) List(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *categoryPostgresRepository) GetByID(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	return &category, err
}
