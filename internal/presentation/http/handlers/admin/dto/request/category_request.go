package request

type CreateCategoryRequest struct {
	Name    string `json:"name" binding:"required"`
	IconURL string `json:"icon_url" binding:"omitempty,url"`
}

type UpdateCategoryRequest struct {
	Name    string `json:"name" binding:"required"`
	IconURL string `json:"icon_url" binding:"omitempty,url"`
}
