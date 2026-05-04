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

func (r *restaurantPostgresRepository) GetByID(ctx context.Context, id uint) (*domain.Restaurant, error) {
	var restaurant domain.Restaurant
	err := r.db.WithContext(ctx).First(&restaurant, id).Error
	return &restaurant, err
}

func (r *restaurantPostgresRepository) GetByPublicID(ctx context.Context, publicID string) (*domain.Restaurant, error) {
	var restaurant domain.Restaurant
	err := r.db.WithContext(ctx).Preload("Images").Where("public_id = ?", publicID).First(&restaurant).Error
	return &restaurant, err
}

func (r *restaurantPostgresRepository) List(ctx context.Context, filter ports.ListRestaurantsFilter) ([]domain.Restaurant, int64, error) {
	var restaurants []domain.Restaurant
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Restaurant{})

	if filter.OwnerID != "" {
		db = db.Where("owner_id = ?", filter.OwnerID)
	}

	if filter.Search != "" {
		db = db.Where("name ILIKE ? OR address ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	// Count total records for pagination
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated data
	err := db.Offset(filter.Offset).Limit(filter.Limit).Order("created_at DESC").Find(&restaurants).Error
	return restaurants, total, err
}

func (r *restaurantPostgresRepository) Update(ctx context.Context, restaurant *domain.Restaurant) error {
	return r.db.WithContext(ctx).Save(restaurant).Error
}

func (r *restaurantPostgresRepository) Delete(ctx context.Context, publicID string) error {
	return r.db.WithContext(ctx).Where("public_id = ?", publicID).Delete(&domain.Restaurant{}).Error
}

func (r *restaurantPostgresRepository) AddImage(ctx context.Context, image *domain.RestaurantImage) error {
	return r.db.WithContext(ctx).Create(image).Error
}
