package request

type UpdateRestaurantRequest struct {
	Name        *string `json:"name" binding:"omitempty,min=8"`
	Address     *string `json:"address" binding:"omitempty,min=8"`
	Description *string `json:"description" binding:"omitempty,min=8"`
	LogoURL     *string `json:"logo_url" binding:"omitempty,url"`
	BannerURL   *string `json:"banner_url" binding:"omitempty,url"`
}
