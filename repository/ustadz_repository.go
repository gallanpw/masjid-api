package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// GetAllUstadz mengambil semua data ustadz dari database
func GetAllUstadz() ([]models.Ustadz, error) {
	var ustadzs []models.Ustadz
	if err := config.DB.Find(&ustadzs).Error; err != nil {
		return nil, err
	}
	return ustadzs, nil
}

// GetUstadzByID mengambil data ustadz berdasarkan ID
func GetUstadzByID(id uint) (*models.Ustadz, error) {
	var ustadz models.Ustadz
	if err := config.DB.First(&ustadz, id).Error; err != nil {
		return nil, err
	}
	return &ustadz, nil
}

// CreateUstadz membuat data ustadz baru
func CreateUstadz(ustadz *models.Ustadz) error {
	return config.DB.Create(ustadz).Error
}

// UpdateUstadz memperbarui data ustadz
func UpdateUstadz(ustadz *models.Ustadz) error {
	return config.DB.Save(ustadz).Error
}

// DeleteUstadz menghapus data ustadz secara lunak
func DeleteUstadz(id uint) error {
	return config.DB.Delete(&models.Ustadz{}, id).Error
}
