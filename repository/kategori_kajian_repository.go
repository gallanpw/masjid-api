package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// GetAllKategoriKajian mengambil semua data kategori kajian
func GetAllKategoriKajian() ([]models.KategoriKajian, error) {
	var kategoriKajians []models.KategoriKajian
	if err := config.DB.Find(&kategoriKajians).Error; err != nil {
		return nil, err
	}
	return kategoriKajians, nil
}

// GetKategoriKajianByID mengambil data kategori kajian berdasarkan ID
func GetKategoriKajianByID(id uint) (*models.KategoriKajian, error) {
	var kategoriKajian models.KategoriKajian
	if err := config.DB.First(&kategoriKajian, id).Error; err != nil {
		return nil, err
	}
	return &kategoriKajian, nil
}

// CreateKategoriKajian membuat data kategori kajian baru
func CreateKategoriKajian(kategoriKajian *models.KategoriKajian) error {
	return config.DB.Create(kategoriKajian).Error
}

// UpdateKategoriKajian memperbarui data kategori kajian
func UpdateKategoriKajian(kategoriKajian *models.KategoriKajian) error {
	return config.DB.Save(kategoriKajian).Error
}

// DeleteKategoriKajian menghapus data kategori kajian secara lunak
func DeleteKategoriKajian(id uint) error {
	return config.DB.Delete(&models.KategoriKajian{}, id).Error
}
