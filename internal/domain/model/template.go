package model

import "time"

type CourseTemplateType string

const (
	TemplateOnline  CourseTemplateType = "online"
	TemplateOffline CourseTemplateType = "offline"
)

type CourseLevel string

const (
	LevelBeginner     CourseLevel = "beginner"
	LevelIntermediate CourseLevel = "intermediate"
	LevelAdvanced     CourseLevel = "advanced"
)

type Template struct {
	ID          int64
	Name        string
	Level       CourseLevel
	Type        CourseTemplateType
	Language    string
	Description *string
	Image       *string
	Price       float64
	Discount    float64
	Duration    *string
	Capacity    int
	Status      string
	CreatedBy   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
