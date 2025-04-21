package handler

import (
	"sue_backend/internal/common/response"
	"sue_backend/internal/domain/service"
	"sue_backend/internal/transport/dto"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: svc,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	user := req.ToUserModel()
	token, err := h.authService.RegisterUser(user)
	if err != nil {
		response.WrapError(c, 400, "Register failed", err.Error())
		return
	}

	response.WrapCreated(c, "User registered successfully", gin.H{
		"token": token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		response.WrapError(c, 401, "Login failed", err.Error())
		return
	}

	response.WrapSuccess(c, "Login success", gin.H{
		"token": token,
	})
}
