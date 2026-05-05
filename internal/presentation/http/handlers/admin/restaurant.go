package admin

import (
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin/dto/request"
	"gopher-restaurant-service/pkg/response"
	"gopher-restaurant-service/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	createUC ports.IAdminCreateRestaurantUseCase
	updateUC ports.IAdminUpdateRestaurantUseCase
	deleteUC ports.IAdminDeleteRestaurantUseCase
	listMyUC ports.IAdminListMyRestaurantsUseCase
}

func NewRestaurantHandler(
	createUC ports.IAdminCreateRestaurantUseCase,
	updateUC ports.IAdminUpdateRestaurantUseCase,
	deleteUC ports.IAdminDeleteRestaurantUseCase,
	listMyUC ports.IAdminListMyRestaurantsUseCase,
) *RestaurantHandler {
	return &RestaurantHandler{
		createUC: createUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		listMyUC: listMyUC,
	}
}

func (h *RestaurantHandler) Create(c *gin.Context) {
	var req request.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldErrors := response.ParseValidationErrors(err)
		response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
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

	restaurant, err := h.createUC.Execute(c.Request.Context(), input)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create restaurant")
		return
	}

	response.Success(c, http.StatusCreated, map[string]interface{}{
		"id": restaurant.ID,
	})
}

func (h *RestaurantHandler) GetMyRestaurants(c *gin.Context) {
	var query request.AdminRestaurantQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		fieldErrors := response.ParseValidationErrors(err)
		response.ValidationError(c, http.StatusBadRequest, "invalid query parameters", fieldErrors)
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

	restaurants, total, err := h.listMyUC.Execute(c.Request.Context(), input)
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
		fieldErrors := response.ParseValidationErrors(err)
		response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
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

	restaurant, err := h.updateUC.Execute(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, domain.ErrRestaurantNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, domain.ErrUnauthorized) {
			response.Error(c, http.StatusForbidden, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, map[string]interface{}{
		"id": restaurant.ID,
	})
}

func (h *RestaurantHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("public_user_id")

	if err := h.deleteUC.Execute(c.Request.Context(), ownerID, id); err != nil {
		if errors.Is(err, domain.ErrRestaurantNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, domain.ErrUnauthorized) {
			response.Error(c, http.StatusForbidden, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "restaurant deleted"})
}
