package request

type UpdateRestaurantRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url" binding:"required,url"`
	BannerURL   string `json:"banner_url" binding:"required,url"`
}
