package route

import (
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCourseTemplateRoutes(rg *gin.RouterGroup, svc *service.CourseTemplateService) {
	h := handler.NewCourseTemplateHandler(svc)
	courseTemplate := rg.Group("/course-templates")
	{
		courseTemplate.POST("", h.Create)
		courseTemplate.GET("/:id", h.Get)
		courseTemplate.GET("", h.List)
		courseTemplate.PUT("/:id", h.Update)
		courseTemplate.PATCH("/:id", h.Update)
		courseTemplate.DELETE("/:id", h.Delete)
	}
}
