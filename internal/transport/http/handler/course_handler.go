package handler

import (
	"net/http"
	"sue_backend/internal/common/response"
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/dto"
	"sue_backend/internal/transport/http/utils"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService *service.CourseService
}

func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: svc}
}

func (h *CourseHandler) Create(c *gin.Context) {
	var req dto.CourseCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}
	course, err := h.courseService.Create(req.ToModel())
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to create course", err.Error())
		return
	}
	response.WrapCreated(c, "Course created successfully", course)
}

func (h *CourseHandler) Get(c *gin.Context) {
	id, exists := dto.ParseIDParam(c)
	if exists != nil {
		response.BadRequest(c, "Invalid course ID", nil)
		return
	}
	course, err := h.courseService.Get(id)
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to get course", err.Error())
		return
	}
	if course == nil {
		response.NotFound(c, "Course not found")
		return
	}
	response.WrapSuccess(c, "Course retrieved successfully", course)
}

func (h *CourseHandler) List(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	courses, total, err := h.courseService.List(page, limit)
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to list courses", err.Error())
		return
	}
	response.WrapSuccess(c, "Courses retrieved successfully", gin.H{"total": total, "courses": courses})
}

func (h *CourseHandler) Update(c *gin.Context) {
	id, exists := dto.ParseIDParam(c)
	if exists != nil {
		response.BadRequest(c, "Invalid course ID", nil)
		return
	}
	var req dto.CourseUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}
	course := req.ToModel()
	course.ID = id
	if err := h.courseService.Update(course); err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to update course", err.Error())
		return
	}
	response.WrapSuccess(c, "Course updated successfully", course)
}

func (h *CourseHandler) Delete(c *gin.Context) {
	id, exists := dto.ParseIDParam(c)
	if exists != nil {
		response.BadRequest(c, "Invalid course ID", nil)
		return
	}
	if err := h.courseService.Delete(id); err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to delete course", err.Error())
		return
	}
	response.WrapSuccess(c, "Course deleted successfully", nil)
}
