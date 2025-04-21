package http

import (
	"sue_backend/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()
	// router.Use(gin.Recovery(), middleware.CORS(), middleware.Logger())

	// api := router.Group("/api/v0")
	// {
	// 	New
	// }

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	// return router
}
