package model

import "time"

type UserRole string

const (
	RoleAdmin   UserRole = "Admin"
	RoleTeacher UserRole = "Teacher"
	RoleStudent UserRole = "Student"
)

type User struct {
	ID              int64     `json:"id"`
	First_name      *string   `json:"first_name" validate:"required,min=3,max=50"`
	Last_name       *string   `json:"last_name" validate:"required,min=3,max=50"`
	Email           *string   `json:"email" validate:"required,email"`
	Password        *string   `json:"password" validate:"required,min=6,max=50"`
	Phone           *string   `json:"phone" validate:"required"`
	Role            UserRole  `json:"role" validate:"required,oneof=Admin Teacher Student"`
	Token           *string   `json:"token"`
	RefreshToken    *string   `json:"refresh_token"`
	RefreshTokenExp time.Time `json:"refresh_token_exp"`
	Status          string    `json:"status" validate:"required,oneof=active inactive deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Salt            *string   `json:"salt"`
}
