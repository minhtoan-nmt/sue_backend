package route

import (
	"sue_backend/internal/common/middleware"
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, userSvc *service.UserService) {
	h := handler.NewUserHandler(userSvc)

	rg.GET("/users/me", h.Me)

	admin := rg.Group("/users")
	admin.Use(middleware.RequireAdmin())
	{
		admin.POST("", h.Create)
		admin.GET("/:id", h.Get)
		admin.GET("", h.List)
		admin.PATCH("/:id", h.Update)
		admin.DELETE("/:id", h.Delete)
	}
}
