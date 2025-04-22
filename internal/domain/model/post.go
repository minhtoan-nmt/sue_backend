package model

import "time"

type PostStatus string

const (
	PostVisible PostStatus = "visible"
	PostHidden  PostStatus = "hidden"
)

type Post struct {
	ID        int64
	ThreadID  int64
	AuthorID  int64
	Content   string
	Status    PostStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
