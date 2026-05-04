package admin

import (
	"gopher-restaurant-service/internal/presentation/http/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	restaurantHandler *RestaurantHandler
	foodHandler       *FoodHandler
	mediaHandler      *MediaHandler
	foodCategoryHandler *FoodCategoryHandler
}

func NewRouter(
	restaurantHandler *RestaurantHandler,
	foodHandler *FoodHandler,
	mediaHandler *MediaHandler,
	foodCategoryHandler *FoodCategoryHandler,
) *Router {
	return &Router{
		restaurantHandler: restaurantHandler,
		foodHandler:       foodHandler,
		mediaHandler:      mediaHandler,
		foodCategoryHandler: foodCategoryHandler,
	}
}

func (r *Router) Register(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.Use(middleware.GatewayAuth())
	{
		// Media management
		admin.POST("/media/presigned-url", r.mediaHandler.GetUploadURL)
		// Restaurant management
		admin.POST("/restaurants", r.restaurantHandler.Create)
		admin.GET("/restaurants", r.restaurantHandler.GetMyRestaurants)
		admin.PUT("/restaurants/:id", r.restaurantHandler.Update)
		admin.DELETE("/restaurants/:id", r.restaurantHandler.Delete)
		
		// Food management
		admin.POST("/foods", r.foodHandler.Create)
		admin.GET("/restaurants/:id/foods", r.foodHandler.GetMenu)
		admin.PUT("/foods/:food_id", r.foodHandler.Update)
		admin.DELETE("/foods/:food_id", r.foodHandler.Delete)

		// Food Category management
		admin.POST("/food-categories", r.foodCategoryHandler.Create)
		admin.PUT("/food-categories/:id", r.foodCategoryHandler.Update)
		admin.DELETE("/food-categories/:id", r.foodCategoryHandler.Delete)
		admin.GET("/food-categories", r.foodCategoryHandler.List)
	}
}
