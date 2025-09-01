package models

import (
	"time"

	"gorm.io/gorm"
)

type Kajian struct {
	// gorm.Model
	ID               uint           `json:"id" gorm:"primaryKey"`
	Title            string         `json:"title" gorm:"not null"`
	Description      string         `json:"description"`
	Date             time.Time      `json:"date" gorm:"not null"`
	UstadzID         uint           `json:"ustadz_id"`
	KategoriKajianID uint           `json:"kategori_kajian_id"`
	Ustadz           Ustadz         `json:"ustadz" gorm:"foreignKey:UstadzID"`
	KategoriKajian   KategoriKajian `json:"kategori" gorm:"foreignKey:KategoriKajianID"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
