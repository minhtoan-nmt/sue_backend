package model

import "time"

type EnrollmentStatus string

const (
	EnrollEnrolled  EnrollmentStatus = "enrolled"
	EnrollCompleted EnrollmentStatus = "completed"
	EnrollDropped   EnrollmentStatus = "dropped"
)

type ActivityType string

const (
	ActivityQuiz       ActivityType = "quiz"
	ActivityAssignment ActivityType = "assignment"
	ActivityForum      ActivityType = "forum"
	ActivityMaterial   ActivityType = "material"
)

type StudentEnrollment struct {
	ID               int64
	StudentID        int64
	CourseID         int64
	Status           EnrollmentStatus
	Progress         int
	LastAccess       *time.Time
	LastActivity     *time.Time
	LastActivityType ActivityType
	EnrolledAt       time.Time
}
