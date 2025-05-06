package model

import "time"

type TeacherAssignmentRole string

const (
	TeacherMain TeacherAssignmentRole = "Teacher"
	TeacherAsst TeacherAssignmentRole = "Assistant"
)

type TeacherAssignmentStatus string

const (
	StatusActive   TeacherAssignmentStatus = "active"
	StatusInactive TeacherAssignmentStatus = "inactive"
	StatusDeleted  TeacherAssignmentStatus = "deleted"
)

type TeacherAssignment struct {
	ID        int64
	TeacherID int64
	CourseID  int64
	Role      TeacherAssignmentRole
	Status    TeacherAssignmentStatus
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
