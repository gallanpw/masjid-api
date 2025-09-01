package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// CreateUser membuat pengguna baru
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetAllUsers mengambil semua pengguna dengan relasi Role
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID mengambil pengguna berdasarkan ID dengan relasi Role
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser memperbarui pengguna
func UpdateUser(user *models.User) error {
	// return config.DB.Save(user).Error
	// Menggunakan Updates() untuk memastikan foreign keys terupdate
	return config.DB.Model(user).Updates(models.User{
		Username: user.Username,
		Password: user.Password,
		RoleID:   user.RoleID,
	}).Error
}

// DeleteUser menghapus pengguna berdasarkan ID
func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
