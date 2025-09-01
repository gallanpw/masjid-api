package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// User merepresentasikan tabel users
type User struct {
	// gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"`
	RoleID    uint           `json:"role_id"`
	Role      Role           `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Claims adalah struct kustom untuk JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
