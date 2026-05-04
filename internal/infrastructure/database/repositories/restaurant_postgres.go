package repositories

import (
	"context"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"

	"gorm.io/gorm"
)

type restaurantPostgresRepository struct {
	db *gorm.DB
}

func NewRestaurantPostgresRepository(db *gorm.DB) ports.RestaurantRepository {
	return &restaurantPostgresRepository{
		db: db,
	}
}

func (r *restaurantPostgresRepository) Create(ctx context.Context, restaurant *domain.Restaurant) error {
	return r.db.WithContext(ctx).Create(restaurant).Error
}

func (r *restaurantPostgresRepository) GetByPublicID(ctx context.Context, publicID string) (*domain.Restaurant, error) {
	var restaurant domain.Restaurant
	err := r.db.WithContext(ctx).Preload("Images").Where("public_id = ?", publicID).First(&restaurant).Error
	return &restaurant, err
}

func (r *restaurantPostgresRepository) ListByOwner(ctx context.Context, ownerID string) ([]domain.Restaurant, error) {
	var restaurants []domain.Restaurant
	err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&restaurants).Error
	return restaurants, err
}

func (r *restaurantPostgresRepository) ListAll(ctx context.Context) ([]domain.Restaurant, error) {
	var restaurants []domain.Restaurant
	err := r.db.WithContext(ctx).Where("status = ?", "active").Find(&restaurants).Error
	return restaurants, err
}

func (r *restaurantPostgresRepository) Update(ctx context.Context, restaurant *domain.Restaurant) error {
	return r.db.WithContext(ctx).Save(restaurant).Error
}

func (r *restaurantPostgresRepository) AddImage(ctx context.Context, image *domain.RestaurantImage) error {
	return r.db.WithContext(ctx).Create(image).Error
}
