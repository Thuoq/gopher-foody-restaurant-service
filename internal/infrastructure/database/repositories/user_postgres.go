package repositories

import (
	"context"

	"gorm.io/gorm"

	"gopher-identity-service/internal/core/domain"
	"gopher-identity-service/internal/core/ports"
)

type userPostgresRepo struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) ports.UserRepository {
	// Auto migrate user schema
	_ = db.AutoMigrate(&domain.User{})

	return &userPostgresRepo{
		db: db,
	}
}

func (r *userPostgresRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
