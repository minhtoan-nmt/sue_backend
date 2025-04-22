package model

import "time"

type FeedbackStatus string

const (
	FeedbackPending  FeedbackStatus = "pending"
	FeedbackApproved FeedbackStatus = "approved"
	FeedbackRejected FeedbackStatus = "rejected"
)

type Feedback struct {
	ID        int64
	TeacherID int64
	StudentID int64
	CourseID  int64
	Comment   string
	Score     float64
	Status    FeedbackStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
