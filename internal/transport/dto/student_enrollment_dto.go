package dto

type EnrollStudentsRequest struct {
	UserIDs []int64 `json:"user_ids" binding:"required,min=1,dive,gt=0"`
}

type StudentEnrollmentResponse struct {
	ID               int64   `json:"id"`
	StudentID        int64   `json:"student_id"`
	CourseID         int64   `json:"course_id"`
	Status           string  `json:"status"`
	Progress         int     `json:"progress"`
	LastAccess       *string `json:"last_access"`
	LastActivity     *string `json:"last_activity"`
	LastActivityType *string `json:"last_activity_type"`
	EnrolledAt       string  `json:"enrolled_at"`
}
