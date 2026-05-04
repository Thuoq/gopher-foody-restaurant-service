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
