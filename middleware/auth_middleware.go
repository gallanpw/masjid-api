package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"masjid-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware memverifikasi JWT token dari header Authorization
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		// Format token: Bearer <token>
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString = parts[1]

		// Parsing token dengan claims yang sudah kita buat
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode penandatanganan sesuai
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired or invalid: " + err.Error()})
		// 	c.Abort()
		// 	return
		// }

		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
				c.Abort()
				return
			}
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired or not yet active"})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan username ke dalam context Gin
		c.Set("username", claims.Username)
		c.Next()
	}
}
