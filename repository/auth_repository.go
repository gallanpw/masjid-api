package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// FindUserByUsername mencari user di database berdasarkan username.
func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
