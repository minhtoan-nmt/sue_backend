package dto

type AssignTeacherRequest struct {
	TeacherID int64  `json:"teacher_id" binding:"required"`
	Role      string `json:"role" binding:"required,oneof=Teacher Assistant"`
}

type TeacherAssignmentResponse struct {
	ID        int64  `json:"id"`
	TeacherID int64  `json:"teacher_id"`
	CourseID  int64  `json:"course_id"`
	Role      string `json:"role"`
}
