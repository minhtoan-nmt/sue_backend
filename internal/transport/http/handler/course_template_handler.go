// internal/transport/handler/course_template_handler.go
package handler

import (
	"net/http"
	"sue_backend/internal/common/response"
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/dto"
	"sue_backend/internal/transport/http/utils"

	"github.com/gin-gonic/gin"
)

type CourseTemplateHandler struct {
	courseTemplateService *service.CourseTemplateService
}

func NewCourseTemplateHandler(svc *service.CourseTemplateService) *CourseTemplateHandler {
	return &CourseTemplateHandler{courseTemplateService: svc}
}
func (h *CourseTemplateHandler) Create(c *gin.Context) {
	var req dto.CourseTemplateCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}
	courseTemplate, err := h.courseTemplateService.Create(req.ToModel())
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to create course template", err.Error())
		return
	}
	response.WrapCreated(c, "Course template created successfully", courseTemplate)
}
func (h *CourseTemplateHandler) Get(c *gin.Context) {
	id, exists := dto.ParseIDParam(c)
	if exists != nil {
		response.BadRequest(c, "Invalid course template ID", nil)
		return
	}
	courseTemplate, err := h.courseTemplateService.Get(id)
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to get course template", err.Error())
		return
	}
	if courseTemplate == nil {
		response.NotFound(c, "Course template not found")
		return
	}
	response.WrapSuccess(c, "Course template retrieved successfully", courseTemplate)
}
func (h *CourseTemplateHandler) List(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	courseTemplates, total, err := h.courseTemplateService.List(page, limit)
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to list course templates", err.Error())
		return
	}
	response.WrapSuccess(c, "Course templates retrieved successfully", gin.H{
		"total":            total,
		"course_templates": courseTemplates,
	})
}
func (h *CourseTemplateHandler) Update(c *gin.Context) {
	id, exists := dto.ParseIDParam(c)
	if exists != nil {
		response.BadRequest(c, "Invalid course template ID", nil)
		return
	}
	var req dto.CourseTemplateUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}
	courseTemplate := req.ToModel()
	courseTemplate.ID = id
	if err := h.courseTemplateService.Update(courseTemplate); err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to update course template", err.Error())
		return
	}
	response.WrapSuccess(c, "Course template updated successfully", courseTemplate)
}
func (h *CourseTemplateHandler) Delete(c *gin.Context) {
	id, exists := dto.ParseIDParam(c)
	if exists != nil {
		response.BadRequest(c, "Invalid course template ID", nil)
		return
	}
	if err := h.courseTemplateService.Delete(id); err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to delete course template", err.Error())
		return
	}
	response.WrapSuccess(c, "Course template deleted successfully", nil)
}
