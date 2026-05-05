package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"go.uber.org/zap"

	"gopher-restaurant-service/internal/application/usecases/food"
	"gopher-restaurant-service/internal/application/usecases/food_category"
	"gopher-restaurant-service/internal/application/usecases/media"
	"gopher-restaurant-service/internal/application/usecases/restaurant"
	"gopher-restaurant-service/internal/config"
	"gopher-restaurant-service/internal/infrastructure/database"
	"gopher-restaurant-service/internal/infrastructure/database/repositories"
	"gopher-restaurant-service/internal/infrastructure/storage"
	httpRouter "gopher-restaurant-service/internal/presentation/http"
	"gopher-restaurant-service/internal/presentation/http/handlers/admin"
	"gopher-restaurant-service/internal/presentation/http/handlers/user"
	"gopher-restaurant-service/pkg/logger"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Core dependencies
	container.Provide(config.LoadConfig)
	container.Provide(logger.NewLogger)

	// Infrastructure
	container.Provide(database.NewPostgresDB)
	container.Provide(repositories.NewRestaurantPostgresRepository)
	container.Provide(repositories.NewFoodPostgresRepository)
	container.Provide(repositories.NewFoodCategoryPostgresRepository)
	container.Provide(storage.NewS3StorageService)

	// Application - Restaurant
	container.Provide(restaurant.NewAdminCreateRestaurantUseCase)
	container.Provide(restaurant.NewAdminUpdateRestaurantUseCase)
	container.Provide(restaurant.NewAdminDeleteRestaurantUseCase)
	container.Provide(restaurant.NewAdminListMyRestaurantsUseCase)
	container.Provide(restaurant.NewUserListRestaurantsUseCase)
	container.Provide(restaurant.NewUserGetRestaurantDetailUseCase)

	// Application - Food
	container.Provide(food.NewAdminCreateFoodUseCase)
	container.Provide(food.NewAdminUpdateFoodUseCase)
	container.Provide(food.NewAdminDeleteFoodUseCase)
	container.Provide(food.NewUserListMenuUseCase)

	// Application - Media
	container.Provide(media.NewGetUploadURLUseCase)

	// Application - Food Category
	container.Provide(food_category.NewAdminCreateFoodCategoryUseCase)
	container.Provide(food_category.NewAdminUpdateFoodCategoryUseCase)
	container.Provide(food_category.NewAdminDeleteFoodCategoryUseCase)
	container.Provide(food_category.NewUserListFoodCategoriesUseCase)

	// Presentation
	container.Provide(admin.NewRestaurantHandler)
	container.Provide(admin.NewFoodHandler)
	container.Provide(admin.NewMediaHandler)
	container.Provide(admin.NewFoodCategoryHandler)
	container.Provide(admin.NewRouter)

	container.Provide(user.NewRestaurantHandler)
	container.Provide(user.NewRouter)

	container.Provide(httpRouter.NewRouter)

	return container
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(cfg *config.Config, log *zap.Logger, router *gin.Engine) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		// 1. Start HTTP Server
		httpServer := &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.App.HTTPPort),
			Handler: router,
		}

		go func() {
			log.Info("Starting Restaurant HTTP Server", zap.Int("port", cfg.App.HTTPPort))
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("HTTP Server failed to start", zap.Error(err))
			}
		}()

		// 2. Wait for interrupt signal to gracefully shutdown
		<-ctx.Done()
		log.Info("Shutting down Restaurant Service gracefully...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// HTTP Shutdown
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Error("HTTP Server forced to shutdown", zap.Error(err))
		}

		log.Info("Service exited gracefully")
	})

	if err != nil {
		fmt.Printf("Error starting application: %v\n", err)
		os.Exit(1)
	}
}
