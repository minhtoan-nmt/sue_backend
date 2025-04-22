package middleware

import (
	"log"
	"net/http"
	"strings"
	"sue_backend/internal/common/response"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.WrapError(c, http.StatusUnauthorized, "Missing or invalid Authorization header", nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			response.WrapError(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			response.WrapError(c, http.StatusUnauthorized, "Invalid token claims", nil)
			c.Abort()
			return
		}
		userID := int64(userIDFloat)
		log.Printf("Role from token: %v", claims["role"])
		roleStr, _ := claims["role"].(string)
		log.Printf("Parsed role: %s", roleStr)
		// inject user_id, role into context if needed
		c.Set("user_id", userID)
		c.Set("role", roleStr)
		c.Next()
	}
}
