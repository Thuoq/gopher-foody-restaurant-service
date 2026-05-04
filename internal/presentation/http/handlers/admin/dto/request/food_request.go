package request

type CreateFoodRequest struct {
	RestaurantPublicID string   `json:"restaurant_id" binding:"required"`
	CategoryID         uint     `json:"category_id"`
	Name               string   `json:"name" binding:"required"`
	Description        string   `json:"description"`
	Price              float64  `json:"price" binding:"required"`
	Quantity           int      `json:"quantity"`
	Images             []string `json:"images"`
}
