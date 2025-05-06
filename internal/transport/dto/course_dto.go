package dto

import (
	"sue_backend/internal/domain/model"
	"time"
)

type CourseCreateRequest struct {
	Name       string    `json:"name" binding:"required"`
	TemplateID int64     `json:"template_id"`
	Schedule   *string   `json:"schedule"`
	Status     string    `json:"status"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

func (c *CourseCreateRequest) ToModel() *model.Course {
	return &model.Course{
		Name:       c.Name,
		TemplateID: c.TemplateID,
		Schedule:   c.Schedule,
		Status:     c.Status,
		StartDate:  &c.StartDate,
		EndDate:    &c.EndDate,
	}
}

type CourseUpdateRequest struct {
	Name      *string    `json:"name"`
	Schedule  *string    `json:"schedule"`
	Status    *string    `json:"status"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

func (c *CourseUpdateRequest) ToModel() *model.Course {
	return &model.Course{
		Name:      *c.Name,
		Schedule:  c.Schedule,
		Status:    *c.Status,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
	}
}
