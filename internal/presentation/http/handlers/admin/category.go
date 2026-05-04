package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	createUC ports.IAdminCreateCategoryUseCase
	updateUC ports.IAdminUpdateCategoryUseCase
	deleteUC ports.IAdminDeleteCategoryUseCase
	listUC   ports.IUserListCategoriesUseCase
}

func NewCategoryHandler(
	createUC ports.IAdminCreateCategoryUseCase,
	updateUC ports.IAdminUpdateCategoryUseCase,
	deleteUC ports.IAdminDeleteCategoryUseCase,
	listUC ports.IUserListCategoriesUseCase,
) *CategoryHandler {
	return &CategoryHandler{
		createUC: createUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		listUC:   listUC,
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	input := ports.CreateCategoryInput{
		Name:    req.Name,
		IconURL: req.IconURL,
	}

	category, err := h.createUC.Execute(c.Request.Context(), input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	input := ports.UpdateCategoryInput{
		ID:      uint(id),
		Name:    req.Name,
		IconURL: req.IconURL,
	}

	category, err := h.updateUC.Execute(c.Request.Context(), input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.deleteUC.Execute(c.Request.Context(), uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "category deleted"})
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.listUC.Execute(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, categories)
}
