package request

type UpdateFoodRequest struct {
	CategoryID  uint     `json:"category_id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" binding:"required,gt=0"`
	Quantity    int      `json:"quantity" binding:"min=0"`
	Status      string   `json:"status" binding:"required,oneof=available out_of_stock"`
	Images      []string `json:"images" binding:"required,min=1"`
}
