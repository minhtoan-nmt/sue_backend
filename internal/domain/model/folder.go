package model

import "time"

type Folder struct {
	ID          int64
	ParentID    *int64
	CourseID    int64
	Name        string
	Description *string
	Order       int
	Status      string
	CreatedAt   time.Time
}
