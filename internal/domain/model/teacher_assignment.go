package model

type TeacherAssignmentRole string

const (
	TeacherMain TeacherAssignmentRole = "Teacher"
	TeacherAsst TeacherAssignmentRole = "Assistant"
)

type TeacherAssignment struct {
	ID        int64
	TeacherID int64
	CourseID  int64
	Role      TeacherAssignmentRole
}
