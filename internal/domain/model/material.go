package model

import "time"

type MaterialType string

const (
	MaterialVideo    MaterialType = "video"
	MaterialAudio    MaterialType = "audio"
	MaterialDocument MaterialType = "document"
	MaterialLink     MaterialType = "link"
)

type Material struct {
	ID          int64
	TeacherID   int64
	CourseID    int64
	FolderID    *int64
	Name        string
	Description *string
	FileURL     string
	Type        MaterialType
	Status      string
	UploadDate  time.Time
}
