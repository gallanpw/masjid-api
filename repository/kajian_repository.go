package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// GetAllKajian mengambil semua data kajian dengan relasi
func GetAllKajian() ([]models.Kajian, error) {
	var kajians []models.Kajian
	if err := config.DB.Preload("Ustadz").Preload("KategoriKajian").Find(&kajians).Error; err != nil {
		return nil, err
	}
	return kajians, nil
}

// GetKajianByID mengambil data kajian berdasarkan ID dengan relasi
func GetKajianByID(id uint) (*models.Kajian, error) {
	var kajian models.Kajian
	if err := config.DB.Preload("Ustadz").Preload("KategoriKajian").First(&kajian, id).Error; err != nil {
		return nil, err
	}
	return &kajian, nil
}

// CreateKajian membuat data kajian baru
func CreateKajian(kajian *models.Kajian) error {
	return config.DB.Create(kajian).Error
}

// UpdateKajian memperbarui data kajian
func UpdateKajian(kajian *models.Kajian) error {
	// Menggunakan Updates() untuk memastikan foreign keys terupdate
	return config.DB.Model(kajian).Updates(models.Kajian{
		Title:            kajian.Title,
		Description:      kajian.Description,
		Date:             kajian.Date,
		UstadzID:         kajian.UstadzID,
		KategoriKajianID: kajian.KategoriKajianID,
	}).Error
}

// DeleteKajian menghapus data kajian secara lunak
func DeleteKajian(id uint) error {
	return config.DB.Delete(&models.Kajian{}, id).Error
}
