package route

import (
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authSvc *service.AuthService) {
	h := handler.NewAuthHandler(authSvc)

	auth := rg.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}
