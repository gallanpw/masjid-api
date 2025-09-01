package models

import (
	"time"

	"gorm.io/gorm"
)

type Ustadz struct {
	// gorm.Model
	// Name string `gorm:"not null"`
	// Bio  string
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Bio       string         `json:"bio"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
