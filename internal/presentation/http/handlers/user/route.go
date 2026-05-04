package user

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	restaurantHandler *RestaurantHandler
}

func NewRouter(restaurantHandler *RestaurantHandler) *Router {
	return &Router{
		restaurantHandler: restaurantHandler,
	}
}

func (r *Router) Register(api *gin.RouterGroup) {
	// Public user routes (No auth required for viewing)
	restaurants := api.Group("/restaurants")
	{
		restaurants.GET("", r.restaurantHandler.List)
		restaurants.GET("/:id", r.restaurantHandler.GetDetail)
		restaurants.GET("/:id/foods", r.restaurantHandler.GetMenu)
	}
}
