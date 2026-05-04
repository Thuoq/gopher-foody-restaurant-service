package request

type CreateFoodCategoryRequest struct {
	Name    string `json:"name" binding:"required"`
	IconURL string `json:"icon_url" binding:"omitempty,url"`
}

type UpdateFoodCategoryRequest struct {
	Name    string `json:"name" binding:"required"`
	IconURL string `json:"icon_url" binding:"omitempty,url"`
}
