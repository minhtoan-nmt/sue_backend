package handler

import (
	"fmt"
	"net/http"
	"sue_backend/internal/common/response"
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/dto"
	"sue_backend/internal/transport/http/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{userService: svc}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}
	user := req.ToModel()
	created, err := h.userService.Create(user)
	if err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}
	response.WrapCreated(c, "User created successfully", created)
}

func (h *UserHandler) Me(c *gin.Context) {
	userIDRaw, exists := c.Get("user_id")
	fmt.Printf("userID raw type = %T\n", userIDRaw)
	if !exists {
		response.WrapError(c, http.StatusUnauthorized, "User ID missing in context", nil)
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		response.WrapError(c, http.StatusUnauthorized, "Invalid user ID format", nil)
		return
	}

	user, err := h.userService.GetByID(userID)
	if err != nil {
		response.InternalServer(c, "Failed to fetch user", err.Error())
		return
	}
	if user == nil {
		response.NotFound(c, "User not found")
		return
	}

	response.WrapSuccess(c, "Fetched current user", dto.UserToResponse(user))
}

func (h *UserHandler) Get(c *gin.Context) {
	id, err := dto.ParseIDParam(c)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}
	user, err := h.userService.GetByID(id)
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}
	response.WrapSuccess(c, "User fetched", user)
}

// func (h *UserHandler) List(c *gin.Context) {
// 	users, err := h.userService.List()
// 	if err != nil {
// 		response.WrapError(c, http.StatusInternalServerError, "Failed to list users", err.Error())
// 		return
// 	}
// 	response.WrapSuccess(c, "Users fetched", users)
// }

func (h *UserHandler) List(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	users, total, err := h.userService.GetPaginatedUsers(page, limit)
	if err != nil {
		response.InternalServer(c, "Failed to fetch paginated users", err.Error())
		return
	}

	response.WrapPagination(c, "Fetched users", dto.UserListToResponse(users), page, limit, total)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := dto.ParseIDParam(c)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}
	var req dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}
	updated := req.ToModel()
	updated.ID = id
	if err := h.userService.Update(updated); err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}
	response.WrapSuccess(c, "User updated", nil)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := dto.ParseIDParam(c)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}
	if err := h.userService.Delete(id); err != nil {
		response.WrapError(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}
	response.WrapSuccess(c, "User deleted", nil)
}
