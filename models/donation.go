package models

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	// gorm.Model
	// Description string  `gorm:"not null"`
	// Amount      float64 `gorm:"not null"`
	ID        uint           `json:"id" gorm:"primaryKey"`
	Amount    float64        `json:"amount" gorm:"not null"`
	Notes     string         `json:"notes"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
