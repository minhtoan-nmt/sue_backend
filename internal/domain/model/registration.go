package model

import "time"

type RegistrationStatus string

const (
	RegistrationPending   RegistrationStatus = "pending"
	RegistrationConfirmed RegistrationStatus = "confirmed"
	RegistrationCancelled RegistrationStatus = "cancelled"
)

type Registration struct {
	ID               int64
	GuestName        string
	Email            string
	Phone            string
	TemplateID       int64
	CourseID         int64
	Status           RegistrationStatus
	RegistrationDate time.Time
	CreatedAt        time.Time
}
