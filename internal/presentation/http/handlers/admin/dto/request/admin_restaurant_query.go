package request

import "gopher-restaurant-service/internal/presentation/http/handlers/common/dto/request"

type AdminRestaurantQuery struct {
	request.PaginationQuery
	Search string `form:"search"`
	Status string `form:"status"`
}
