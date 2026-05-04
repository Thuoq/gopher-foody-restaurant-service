package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	restaurantUseCase ports.IRestaurantUseCase
}

func NewRestaurantHandler(restaurantUseCase ports.IRestaurantUseCase) *RestaurantHandler {
	return &RestaurantHandler{
		restaurantUseCase: restaurantUseCase,
	}
}

func (h *RestaurantHandler) Create(c *gin.Context) {
	var req request.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get OwnerID from Gateway header (via GatewayAuth middleware)
	ownerID := c.GetString("public_user_id")
	if ownerID == "" {
		response.Error(c, http.StatusUnauthorized, "missing user identity")
		return
	}

	input := ports.CreateRestaurantInput{
		OwnerID:     ownerID,
		Name:        req.Name,
		Address:     req.Address,
		Description: req.Description,
		LogoURL:     req.LogoURL,
		BannerURL:   req.BannerURL,
	}

	restaurant, err := h.restaurantUseCase.CreateRestaurant(c.Request.Context(), input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create restaurant")
		return
	}

	response.Success(c, http.StatusCreated, restaurant)
}

func (h *RestaurantHandler) GetMyRestaurants(c *gin.Context) {
	ownerID := c.GetString("public_user_id")
	restaurants, err := h.restaurantUseCase.GetMyRestaurants(c.Request.Context(), ownerID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch restaurants")
		return
	}

	response.Success(c, http.StatusOK, restaurants)
}
