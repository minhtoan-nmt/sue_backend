package dto

type CreateFolderRequest struct {
	ParentID    *int64  `json:"parent_id"`
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Order       *int    `json:"order"`
	Status      *string `json:"status"`
}

type FolderResponse struct {
	ID          int64   `json:"id"`
	ParentID    *int64  `json:"parent_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Order       int     `json:"order"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
}
