package domain

import "time"

type Food struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	PublicID     string      `json:"public_id" gorm:"unique;not null"`
	RestaurantID uint        `json:"restaurant_id" gorm:"not null"`
	CategoryID   uint        `json:"category_id"`
	Name         string      `json:"name" gorm:"not null"`
	Description  string      `json:"description"`
	Price        float64     `json:"price" gorm:"not null"`
	Quantity     int         `json:"quantity" gorm:"default:0"`
	Status       string      `json:"status" gorm:"default:available"`
	Images       []FoodImage `json:"images" gorm:"foreignKey:FoodID"`
	Category     FoodCategory `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type FoodImage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FoodID    uint      `json:"food_id" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	IsPrimary bool      `json:"is_primary" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}
