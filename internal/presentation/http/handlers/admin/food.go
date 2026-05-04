package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	foodUseCase ports.IFoodUseCase
}

func NewFoodHandler(foodUseCase ports.IFoodUseCase) *FoodHandler {
	return &FoodHandler{
		foodUseCase: foodUseCase,
	}
}

func (h *FoodHandler) Create(c *gin.Context) {
	var req request.CreateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ownerID := c.GetString("public_user_id")

	input := ports.CreateFoodInput{
		RestaurantPublicID: req.RestaurantPublicID,
		CategoryID:         req.CategoryID,
		Name:               req.Name,
		Description:        req.Description,
		Price:              req.Price,
		Quantity:           req.Quantity,
		Images:             req.Images,
	}

	food, err := h.foodUseCase.CreateFood(c.Request.Context(), ownerID, input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, food)
}

func (h *FoodHandler) GetMenu(c *gin.Context) {
	restaurantID := c.Param("id")
	foods, err := h.foodUseCase.ListRestaurantMenu(c.Request.Context(), restaurantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch menu")
		return
	}

	response.Success(c, http.StatusOK, foods)
}
