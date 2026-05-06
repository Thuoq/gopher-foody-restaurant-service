package user

import (
	"errors"
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/user/dto/request"
	dto "gopher-restaurant-service/internal/presentation/http/handlers/user/dto/response"
	"gopher-restaurant-service/pkg/app_response"
	"gopher-restaurant-service/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	listUC   ports.IUserListRestaurantsUseCase
	detailUC ports.IUserGetRestaurantDetailUseCase
	menuUC   ports.IUserListMenuUseCase
}

func NewRestaurantHandler(
	listUC ports.IUserListRestaurantsUseCase,
	detailUC ports.IUserGetRestaurantDetailUseCase,
	menuUC ports.IUserListMenuUseCase,
) *RestaurantHandler {
	return &RestaurantHandler{
		listUC:   listUC,
		detailUC: detailUC,
		menuUC:   menuUC,
	}
}

func (h *RestaurantHandler) List(c *gin.Context) {
	var query request.UserRestaurantQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		fieldErrors := app_response.ParseValidationErrors(err)
		app_response.ValidationError(c, http.StatusBadRequest, "invalid query parameters", fieldErrors)
		return
	}

	input := ports.ListRestaurantsInput{
		Search: query.Search,
		Status: "active", // Customers only see active restaurants
		Page:   query.Page,
		Limit:  query.Limit,
	}

	restaurants, total, err := h.listUC.Execute(c.Request.Context(), input)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, "failed to fetch restaurants")
		return
	}

	resData := make([]dto.RestaurantResponse, len(restaurants))
	for i, r := range restaurants {
		resData[i] = mapRestaurantToDTO(r)
	}

	res := utils.NewPaginatedResponse(resData, total, query.Page, query.Limit)
	app_response.Success(c, http.StatusOK, res)
}

func (h *RestaurantHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	restaurant, err := h.detailUC.Execute(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrRestaurantNotFound) {
			app_response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		app_response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	app_response.Success(c, http.StatusOK, mapRestaurantToDTO(*restaurant))
}

func (h *RestaurantHandler) GetMenu(c *gin.Context) {
	id := c.Param("id")
	foods, err := h.menuUC.Execute(c.Request.Context(), id)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, "failed to fetch menu")
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

	app_response.Success(c, http.StatusOK, res)
}

func mapRestaurantToDTO(r domain.Restaurant) dto.RestaurantResponse {
	foods := make([]dto.FoodResponse, len(r.Foods))
	for i, f := range r.Foods {
		foods[i] = dto.FoodResponse{
			PublicID:    f.PublicID,
			Name:        f.Name,
			Description: f.Description,
			Price:       f.Price,
			Quantity:    f.Quantity,
			Status:      f.Status,
			Category:    f.Category,
			Images:      f.Images,
		}
	}

	return dto.RestaurantResponse{
		PublicID:    r.PublicID,
		Name:        r.Name,
		Address:     r.Address,
		Description: r.Description,
		LogoURL:     r.LogoURL,
		BannerURL:   r.BannerURL,
		Status:      r.Status,
		Images:      r.Images,
		Foods:       foods,
	}
}
