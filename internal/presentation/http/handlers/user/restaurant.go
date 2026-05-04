package user

import (
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	dto "gopher-restaurant-service/internal/presentation/http/handlers/user/dto/response"
	"gopher-restaurant-service/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	restaurantUseCase ports.IRestaurantUseCase
	foodUseCase       ports.IFoodUseCase
}

func NewRestaurantHandler(restaurantUseCase ports.IRestaurantUseCase, foodUseCase ports.IFoodUseCase) *RestaurantHandler {
	return &RestaurantHandler{
		restaurantUseCase: restaurantUseCase,
		foodUseCase:       foodUseCase,
	}
}

func (h *RestaurantHandler) List(c *gin.Context) {
	restaurants, err := h.restaurantUseCase.ListRestaurants(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch restaurants")
		return
	}

	res := make([]dto.RestaurantResponse, len(restaurants))
	for i, r := range restaurants {
		res[i] = mapRestaurantToDTO(r)
	}

	response.Success(c, http.StatusOK, res)
}

func (h *RestaurantHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	restaurant, err := h.restaurantUseCase.GetRestaurantDetail(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "restaurant not found")
		return
	}

	response.Success(c, http.StatusOK, mapRestaurantToDTO(*restaurant))
}

func (h *RestaurantHandler) GetMenu(c *gin.Context) {
	id := c.Param("id")
	foods, err := h.foodUseCase.ListRestaurantMenu(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch menu")
		return
	}

	res := make([]dto.FoodResponse, len(foods))
	for i, f := range foods {
		res[i] = dto.FoodResponse{
			PublicID:    f.PublicID,
			Name:        f.Name,
			Description: f.Description,
			Price:       f.Price,
			Quantity:    f.Quantity,
			Status:      f.Status,
			Images:      f.Images,
		}
	}

	response.Success(c, http.StatusOK, res)
}

func mapRestaurantToDTO(r domain.Restaurant) dto.RestaurantResponse {
	return dto.RestaurantResponse{
		PublicID:    r.PublicID,
		Name:        r.Name,
		Address:     r.Address,
		Description: r.Description,
		LogoURL:     r.LogoURL,
		BannerURL:   r.BannerURL,
		Status:      r.Status,
		Images:      r.Images,
	}
}
