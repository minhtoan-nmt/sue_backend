package model

import "time"

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
)

type Payment struct {
	ID             int64
	RegistrationID int64
	Amount         float64
	Status         PaymentStatus
	PaymentDate    time.Time
	PaymentMethod  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
