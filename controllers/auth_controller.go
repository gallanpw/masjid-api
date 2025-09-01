package controllers

import (
	"net/http"
	"time"

	"masjid-api/models"
	"masjid-api/repository"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// LoginInput adalah struct untuk input login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUser mengautentikasi user dan mengembalikan JWT token
func LoginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.FindUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	// Buat JWT token
	expirationTime := time.Now().Add(3 * time.Hour) // masa berlaku token selama 3 jam
	claims := &models.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// LogoutUser menangani permintaan logout
func LogoutUser(c *gin.Context) {
	// Tidak ada logika khusus di sisi server untuk JWT.
	// Klien bertanggung jawab untuk menghapus token.
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil logout. Silakan hapus token Anda di sisi klien."})
}
