package dto

import "gopher-restaurant-service/internal/core/domain"

type RestaurantResponse struct {
	PublicID    string                   `json:"id"`
	Name        string                   `json:"name"`
	Address     string                   `json:"address"`
	Description string                   `json:"description"`
	LogoURL     string                   `json:"logo_url"`
	BannerURL   string                   `json:"banner_url"`
	Status      string                   `json:"status"`
	Images      []domain.RestaurantImage `json:"images,omitempty"`
}

type FoodResponse struct {
	PublicID    string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Quantity    int                `json:"quantity"`
	Status      string             `json:"status"`
	Images      []domain.FoodImage `json:"images,omitempty"`
}
