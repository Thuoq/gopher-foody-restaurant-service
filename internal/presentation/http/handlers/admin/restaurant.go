package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/response"
	"gopher-restaurant-service/pkg/utils"
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
	var query request.AdminRestaurantQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ownerID := c.GetString("public_user_id")

	input := ports.ListRestaurantsInput{
		OwnerID: ownerID,
		Search:  query.Search,
		Status:  query.Status,
		Page:    query.Page,
		Limit:   query.Limit,
	}

	restaurants, total, err := h.restaurantUseCase.ListRestaurants(c.Request.Context(), input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch restaurants")
		return
	}

	res := utils.NewPaginatedResponse(restaurants, total, query.Page, query.Limit)
	response.Success(c, http.StatusOK, res)
}

func (h *RestaurantHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ownerID := c.GetString("public_user_id")

	input := ports.UpdateRestaurantInput{
		PublicID:    id,
		OwnerID:     ownerID,
		Name:        req.Name,
		Address:     req.Address,
		Description: req.Description,
		LogoURL:     req.LogoURL,
		BannerURL:   req.BannerURL,
	}

	restaurant, err := h.restaurantUseCase.UpdateRestaurant(c.Request.Context(), input)
	if err != nil {
		response.Error(c, http.StatusForbidden, err.Error())
		return
	}

	response.Success(c, http.StatusOK, restaurant)
}

func (h *RestaurantHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("public_user_id")

	if err := h.restaurantUseCase.DeleteRestaurant(c.Request.Context(), ownerID, id); err != nil {
		response.Error(c, http.StatusForbidden, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "restaurant deleted"})
}
