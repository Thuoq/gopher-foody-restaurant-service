package request

type CreateFoodRequest struct {
	RestaurantPublicID string   `json:"restaurant_id" binding:"required"`
	CategoryID         uint     `json:"category_id" binding:"required"`
	Name               string   `json:"name" binding:"required"`
	Description        string   `json:"description"`
	Price              float64  `json:"price" binding:"required,gt=0"`
	Quantity           int      `json:"quantity" binding:"min=0"`
	Images             []string `json:"images" binding:"required,min=1"`
}
