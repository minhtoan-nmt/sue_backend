// internal/transport/dto/course_template_dto.go
package dto

import "sue_backend/internal/domain/model"

type CourseTemplateCreateRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Type        string   `json:"type" binding:"required"`
	Level       string   `json:"level" binding:"required"`
	Language    string   `json:"language" binding:"required"`
	Price       *float64 `json:"price"`
	Discount    *float64 `json:"discount"`
	Duration    *string  `json:"duration"`
	Capacity    *int     `json:"capacity"`
}

func (c *CourseTemplateCreateRequest) ToModel() *model.CourseTemplate {
	return &model.CourseTemplate{
		Name:        c.Name,
		Description: c.Description,
		Type:        c.Type,
		Level:       c.Level,
		Language:    c.Language,
		Price:       c.Price,
		Discount:    c.Discount,
		Duration:    c.Duration,
		Capacity:    c.Capacity,
	}
}

type CourseTemplateUpdateRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Type        string   `json:"type" binding:"required"`
	Level       string   `json:"level" binding:"required"`
	Language    string   `json:"language" binding:"required"`
	Price       *float64 `json:"price"`
	Discount    *float64 `json:"discount"`
	Duration    *string  `json:"duration"`
	Capacity    *int     `json:"capacity"`
}

func (c *CourseTemplateUpdateRequest) ToModel() *model.CourseTemplate {
	return &model.CourseTemplate{
		Name:        c.Name,
		Description: c.Description,
		Type:        c.Type,
		Level:       c.Level,
		Language:    c.Language,
		Price:       c.Price,
		Discount:    c.Discount,
		Duration:    c.Duration,
		Capacity:    c.Capacity,
	}
}

type CourseTemplateResponse struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Status      string   `json:"status"`
	CreatedBy   *int64   `json:"created_by"`
	Type        string   `json:"type"`
	Level       string   `json:"level"`
	Language    string   `json:"language"`
	Image       *string  `json:"image"`
	Price       *float64 `json:"price"`
	Discount    *float64 `json:"discount"`
	Duration    *string  `json:"duration"`
	Capacity    *int     `json:"capacity"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
