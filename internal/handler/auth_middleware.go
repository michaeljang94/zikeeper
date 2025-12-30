package handler

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mxwI8s8mlZo7NSgkvKBBO4BOjFXV4IT8")

func AuthMiddleware() gin.HandlerFunc {
	defaultRoles := []string{"user", "dealer", "admin"}
	return AuthMiddlewareWithRoles(defaultRoles)
}

func AuthMiddlewareWithRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if !slices.Contains(roles, claims["role"].(string)) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
			}

			c.Set("username", claims["username"])
			c.Set("role", claims["role"])

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
