package admin

import (
	"gopher-restaurant-service/internal/presentation/http/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	restaurantHandler *RestaurantHandler
	foodHandler       *FoodHandler
}

func NewRouter(restaurantHandler *RestaurantHandler, foodHandler *FoodHandler) *Router {
	return &Router{
		restaurantHandler: restaurantHandler,
		foodHandler:       foodHandler,
	}
}

func (r *Router) Register(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.Use(middleware.GatewayAuth())
	{
		// Restaurant management
		admin.POST("/restaurants", r.restaurantHandler.Create)
		admin.GET("/restaurants", r.restaurantHandler.GetMyRestaurants)

		// Food management
		admin.POST("/foods", r.foodHandler.Create)
		admin.GET("/restaurants/:id/foods", r.foodHandler.GetMenu)
	}
}
