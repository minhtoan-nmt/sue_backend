package dto

import "sue_backend/internal/domain/model"

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

func (r *RegisterRequest) ToUserModel() *model.User {
	return &model.User{
		First_name: &r.FirstName,
		Last_name:  &r.LastName,
		Email:      &r.Email,
		Password:   &r.Password,
		Phone:      &r.Phone,
		Role:       model.UserRole(r.Role),
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
