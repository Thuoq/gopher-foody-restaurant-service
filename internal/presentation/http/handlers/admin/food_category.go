package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/app_response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodCategoryHandler struct {
	createUC ports.IAdminCreateFoodCategoryUseCase
	updateUC ports.IAdminUpdateFoodCategoryUseCase
	deleteUC ports.IAdminDeleteFoodCategoryUseCase
	listUC   ports.IUserListFoodCategoriesUseCase
}

func NewFoodCategoryHandler(
	createUC ports.IAdminCreateFoodCategoryUseCase,
	updateUC ports.IAdminUpdateFoodCategoryUseCase,
	deleteUC ports.IAdminDeleteFoodCategoryUseCase,
	listUC ports.IUserListFoodCategoriesUseCase,
) *FoodCategoryHandler {
	return &FoodCategoryHandler{
		createUC: createUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		listUC:   listUC,
	}
}

func (h *FoodCategoryHandler) Create(c *gin.Context) {
	var req request.CreateFoodCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldErrors := app_response.ParseValidationErrors(err)
		app_response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
		return
	}

	input := ports.CreateFoodCategoryInput{
		Name:    req.Name,
		IconURL: req.IconURL,
	}

	category, err := h.createUC.Execute(c.Request.Context(), input)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	app_response.Success(c, http.StatusCreated, category)
}

func (h *FoodCategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req request.UpdateFoodCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldErrors := app_response.ParseValidationErrors(err)
		app_response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
		return
	}

	input := ports.UpdateFoodCategoryInput{
		ID:      uint(id),
		Name:    req.Name,
		IconURL: req.IconURL,
	}

	category, err := h.updateUC.Execute(c.Request.Context(), input)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	app_response.Success(c, http.StatusOK, category)
}

func (h *FoodCategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.deleteUC.Execute(c.Request.Context(), uint(id)); err != nil {
		app_response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	app_response.Success(c, http.StatusOK, gin.H{"message": "food category deleted"})
}

func (h *FoodCategoryHandler) List(c *gin.Context) {
	categories, err := h.listUC.Execute(c.Request.Context())
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	app_response.Success(c, http.StatusOK, categories)
}
