package domain

import "time"

type Restaurant struct {
	ID          uint              `json:"id" gorm:"primaryKey"`
	PublicID    string            `json:"public_id" gorm:"unique;not null"`
	OwnerID     string            `json:"owner_id" gorm:"not null"`
	Name        string            `json:"name" gorm:"not null"`
	Address     string            `json:"address" gorm:"not null"`
	Description string            `json:"description"`
	LogoURL     string            `json:"logo_url"`
	BannerURL   string            `json:"banner_url"`
	Status      string            `json:"status" gorm:"default:active"`
	Images      []RestaurantImage `json:"images" gorm:"foreignKey:RestaurantID"`
	Foods       []Food            `json:"foods" gorm:"foreignKey:RestaurantID"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type RestaurantImage struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	RestaurantID uint      `json:"restaurant_id" gorm:"not null"`
	ImageURL     string    `json:"image_url" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
}
