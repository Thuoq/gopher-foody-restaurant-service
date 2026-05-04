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

func (h *FoodHandler) Update(c *gin.Context) {
	id := c.Param("food_id")
	var req request.UpdateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ownerID := c.GetString("public_user_id")

	input := ports.UpdateFoodInput{
		PublicID:    id,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Status:      req.Status,
		Images:      req.Images,
	}

	food, err := h.foodUseCase.UpdateFood(c.Request.Context(), ownerID, input)
	if err != nil {
		response.Error(c, http.StatusForbidden, err.Error())
		return
	}

	response.Success(c, http.StatusOK, food)
}

func (h *FoodHandler) Delete(c *gin.Context) {
	id := c.Param("food_id")
	ownerID := c.GetString("public_user_id")

	if err := h.foodUseCase.DeleteFood(c.Request.Context(), ownerID, id); err != nil {
		response.Error(c, http.StatusForbidden, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "food deleted"})
}
