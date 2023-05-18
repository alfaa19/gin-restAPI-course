package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk memeriksa token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization not found"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		secretKey := []byte("secret-wawawawawaaw")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || (role != "admin" && role != "user") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role is not valid", "role": role})
			c.Abort()
			return
		}

		c.Set("role", role)

		c.Next()
	}
}

// AdminMiddleware adalah middleware untuk memeriksa role admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only Admin can access"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// UserMiddleware adalah middleware untuk memeriksa role user
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" && role != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only User can acces"})
			c.Abort()
			return
		}

		c.Next()
	}
}
