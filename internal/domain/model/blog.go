package model

import "time"

type BlogStatus string

const (
	BlogDraft     BlogStatus = "draft"
	BlogPublished BlogStatus = "published"
	BlogArchived  BlogStatus = "archived"
)

type Blog struct {
	ID              int64
	Title           string
	Content         string
	AuthorID        int64
	Status          BlogStatus
	Tags            *string
	ImageURL        *string
	CommentsCount   int
	LikesCount      int
	ViewsCount      int
	CommentsEnabled bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
