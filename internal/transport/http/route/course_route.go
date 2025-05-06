package route

import (
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCourseRoutes(rg *gin.RouterGroup, svc *service.CourseService) {
	h := handler.NewCourseHandler(svc)
	course := rg.Group("/courses")
	{
		course.POST("", h.Create)
		course.GET("/:id", h.Get)
		course.GET("", h.List)
		course.PATCH("/:id", h.Update)
		course.DELETE("/:id", h.Delete)
	}
}
