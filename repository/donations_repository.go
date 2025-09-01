package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// CreateDonation membuat data donasi baru
func CreateDonation(donation *models.Donation) error {
	return config.DB.Create(donation).Error
}
