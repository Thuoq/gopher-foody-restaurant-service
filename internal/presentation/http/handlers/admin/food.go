package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/app_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	createUC   ports.IAdminCreateFoodUseCase
	updateUC   ports.IAdminUpdateFoodUseCase
	deleteUC   ports.IAdminDeleteFoodUseCase
	listMenuUC ports.IUserListMenuUseCase
}

func NewFoodHandler(
	createUC ports.IAdminCreateFoodUseCase,
	updateUC ports.IAdminUpdateFoodUseCase,
	deleteUC ports.IAdminDeleteFoodUseCase,
	listMenuUC ports.IUserListMenuUseCase,
) *FoodHandler {
	return &FoodHandler{
		createUC:   createUC,
		updateUC:   updateUC,
		deleteUC:   deleteUC,
		listMenuUC: listMenuUC,
	}
}

func (h *FoodHandler) Create(c *gin.Context) {
	var req request.CreateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldErrors := app_response.ParseValidationErrors(err)
		app_response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
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

	food, err := h.createUC.Execute(c.Request.Context(), ownerID, input)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	app_response.Success(c, http.StatusCreated, food)
}

func (h *FoodHandler) GetMenu(c *gin.Context) {
	restaurantID := c.Param("id")
	foods, err := h.listMenuUC.Execute(c.Request.Context(), restaurantID)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, "failed to fetch menu")
		return
	}

	app_response.Success(c, http.StatusOK, foods)
}

func (h *FoodHandler) Update(c *gin.Context) {
	id := c.Param("food_id")
	var req request.UpdateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldErrors := app_response.ParseValidationErrors(err)
		app_response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
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

	food, err := h.updateUC.Execute(c.Request.Context(), ownerID, input)
	if err != nil {
		app_response.Error(c, http.StatusForbidden, err.Error())
		return
	}

	app_response.Success(c, http.StatusOK, food)
}

func (h *FoodHandler) Delete(c *gin.Context) {
	id := c.Param("food_id")
	ownerID := c.GetString("public_user_id")

	if err := h.deleteUC.Execute(c.Request.Context(), ownerID, id); err != nil {
		app_response.Error(c, http.StatusForbidden, err.Error())
		return
	}

	app_response.Success(c, http.StatusOK, gin.H{"message": "food deleted"})
}
