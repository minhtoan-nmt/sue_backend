package middleware

import (
	"net/http"
	"sue_backend/internal/common/response"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	roleSet := make(map[string]struct{}, len(allowedRoles))
	for _, r := range allowedRoles {
		roleSet[r] = struct{}{}
	}

	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			response.WrapError(c, http.StatusUnauthorized, "Missing role in context", nil)
			c.Abort()
			return
		}

		roleStr, ok := roleVal.(string)
		if !ok {
			response.WrapError(c, http.StatusForbidden, "Invalid role format", nil)
			c.Abort()
			return
		}

		if _, allowed := roleSet[roleStr]; !allowed {
			response.WrapError(c, http.StatusForbidden, "Permission denied", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return RequireRole("Admin")
}

func RequireTeacher() gin.HandlerFunc {
	return RequireRole("Teacher")
}

func RequireStudent() gin.HandlerFunc {
	return RequireRole("Student")
}
