package user

import (
	"gopher-restaurant-service/internal/core/domain"
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/internal/presentation/http/handlers/user/dto/request"
	dto "gopher-restaurant-service/internal/presentation/http/handlers/user/dto/response"
	"gopher-restaurant-service/pkg/response"
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
		response.Error(c, http.StatusBadRequest, err.Error())
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
		response.Error(c, http.StatusInternalServerError, "failed to fetch restaurants")
		return
	}

	resData := make([]dto.RestaurantResponse, len(restaurants))
	for i, r := range restaurants {
		resData[i] = mapRestaurantToDTO(r)
	}

	res := utils.NewPaginatedResponse(resData, total, query.Page, query.Limit)
	response.Success(c, http.StatusOK, res)
}

func (h *RestaurantHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	restaurant, err := h.detailUC.Execute(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "restaurant not found")
		return
	}

	response.Success(c, http.StatusOK, mapRestaurantToDTO(*restaurant))
}

func (h *RestaurantHandler) GetMenu(c *gin.Context) {
	id := c.Param("id")
	foods, err := h.menuUC.Execute(c.Request.Context(), id)
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
