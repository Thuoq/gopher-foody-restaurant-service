package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gopher-identity-service/internal/config"
	"gopher-identity-service/internal/presentation/http/handlers/user"
)

func NewRouter(cfg *config.Config, logger *zap.Logger, userRouter *user.Router) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Add basic middlewares
	r.Use(gin.Recovery())

	// Example health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api/v1")
	userRouter.Register(api)

	return r
}
