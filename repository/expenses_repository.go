package repository

import (
	"masjid-api/config"
	"masjid-api/models"
)

// CreateExpense membuat data pengeluaran baru
func CreateExpense(expense *models.Expense) error {
	return config.DB.Create(expense).Error
}
