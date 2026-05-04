package domain

import "time"

type FoodCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	IconURL   string    `json:"icon_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (FoodCategory) TableName() string {
	return "food_categories"
}
