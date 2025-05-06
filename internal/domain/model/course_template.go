// internal/domain/model/course_template.go
package model

import "time"

type CourseTemplate struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
	CreatedBy   *int64    `json:"created_by"`
	Type        string    `json:"type"`
	Level       string    `json:"level"`
	Language    string    `json:"language"`
	Image       *string   `json:"image"`
	Price       *float64  `json:"price"`
	Discount    *float64  `json:"discount"`
	Duration    *string   `json:"duration"`
	Capacity    *int      `json:"capacity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
