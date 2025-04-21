package http

import (
	"sue_backend/config"

	"sue_backend/internal/common/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORS(), middleware.RequestID())

	authMiddleware := middleware.JWTAuth(cfg.Auth.JWTSecret)

	// Public routes
	api := router.Group("/api/v0")
	{
		api.Use(authMiddleware)

	}

	// Private routes
	// protected := api.Group("")
	// protected.Use(authMiddleware)
	// {

	// 	// route.RegisterUserRoutes(protected, cfg)
	// 	// route.RegisterCourseRoutes(protected, cfg)
	// }

	// Healthcheck
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	return router
}
