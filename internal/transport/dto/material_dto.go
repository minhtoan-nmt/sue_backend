package dto

type UploadMaterialRequest struct {
	TeacherID   int64   `json:"teacher_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	FileURL     string  `json:"file_url" binding:"required"`
	FolderID    *int64  `json:"folder_id"`
	Type        *string `json:"type"`
	Status      *string `json:"status"`
}

type MaterialResponse struct {
	ID          int64   `json:"id"`
	TeacherID   int64   `json:"teacher_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	FileURL     string  `json:"file_url"`
	FolderID    *int64  `json:"folder_id"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	UploadDate  string  `json:"upload_date"`
}
