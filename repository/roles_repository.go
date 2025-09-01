package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// CreateRole membuat peran baru
func CreateRole(role *models.Role) error {
	return config.DB.Create(role).Error
}

// GetAllRoles mengambil semua peran
func GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// GetRoleByID mengambil peran berdasarkan ID
func GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// UpdateRole memperbarui peran
func UpdateRole(role *models.Role) error {
	return config.DB.Save(role).Error
}

// DeleteRole menghapus peran berdasarkan ID
func DeleteRole(id uint) error {
	return config.DB.Delete(&models.Role{}, id).Error
}
