package model

import "time"

type ThreadStatus string

const (
	ThreadOpen   ThreadStatus = "open"
	ThreadClosed ThreadStatus = "closed"
)

type Thread struct {
	ID        int64
	ForumID   int64
	Title     string
	Content   string
	Status    ThreadStatus
	Views     int
	Replies   int
	LastReply *time.Time
	CreatedBy int64
	CreatedAt time.Time
}
