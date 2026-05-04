package request

type CreateRestaurantRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	BannerURL   string `json:"banner_url"`
}
