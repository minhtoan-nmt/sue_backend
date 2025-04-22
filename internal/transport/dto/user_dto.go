package dto

import (
	"errors"
	"strconv"
	"time"

	"sue_backend/internal/domain/model"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=50"`
	LastName  string `json:"last_name" binding:"required,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,max=50"`
	Phone     string `json:"phone" binding:"required"`
	Role      string `json:"role" binding:"required,oneof=Admin Teacher Student"`
}

func (r *UserCreateRequest) ToModel() *model.User {
	return &model.User{
		First_name: &r.FirstName,
		Last_name:  &r.LastName,
		Email:      &r.Email,
		Password:   &r.Password,
		Phone:      &r.Phone,
		Role:       model.UserRole(r.Role),
		Status:     "active",
	}
}

type UserUpdateRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	Status    *string `json:"status" binding:"omitempty,oneof=active inactive deleted"`
}

func (r *UserUpdateRequest) ToModel() *model.User {
	return &model.User{
		First_name: r.FirstName,
		Last_name:  r.LastName,
		Phone:      r.Phone,
		Status:     getStrValueOrDefault(r.Status, "active"),
	}
}

type UserResponse struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserToResponse(u *model.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		FirstName: deref(u.First_name),
		LastName:  deref(u.Last_name),
		Email:     deref(u.Email),
		Phone:     deref(u.Phone),
		Role:      string(u.Role),
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserListToResponse(users []*model.User) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, *UserToResponse(u))
	}
	return res
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func getStrValueOrDefault(ptr *string, fallback string) string {
	if ptr != nil {
		return *ptr
	}
	return fallback
}

func ParseIDParam(c *gin.Context) (int64, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid ID")
	}
	return id, nil
}
