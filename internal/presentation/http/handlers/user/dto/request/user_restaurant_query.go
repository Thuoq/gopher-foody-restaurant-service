package request

import "gopher-restaurant-service/internal/presentation/http/handlers/common/dto/request"

type UserRestaurantQuery struct {
	request.PaginationQuery
	Search string `form:"search"`
	Status string `form:"status"`
}
