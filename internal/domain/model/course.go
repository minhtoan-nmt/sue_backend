package model

import "time"

type Course struct {
	ID         int64
	Name       string
	TemplateID int64
	Schedule   *string
	Status     string
	StartDate  *time.Time
	EndDate    *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	TeacherAssignments []*TeacherAssignment `json:"teacher_assignments"`
}
