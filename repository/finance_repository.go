package repository

import (
	"masjid-api/config"
	"masjid-api/models"
	"time"
)

// GetAllDonations mengambil semua data donasi
func GetAllDonations() ([]models.Donation, error) {
	var donations []models.Donation
	if err := config.DB.Find(&donations).Error; err != nil {
		return nil, err
	}
	return donations, nil
}

// GetAllExpenses mengambil semua data pengeluaran
func GetAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	if err := config.DB.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

// GetAllDonationsInDateRange mengambil donasi dalam rentang tanggal tertentu
func GetAllDonationsInDateRange(start, end time.Time) ([]models.Donation, error) {
	var donations []models.Donation
	if err := config.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&donations).Error; err != nil {
		return nil, err
	}
	return donations, nil
}

// GetAllExpensesInDateRange mengambil pengeluaran dalam rentang tanggal tertentu
func GetAllExpensesInDateRange(start, end time.Time) ([]models.Expense, error) {
	var expenses []models.Expense
	if err := config.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

// CreateFinanceEntry membuat entri baru di tabel finance
func CreateFinanceEntry(finance *models.Finance) error {
	return config.DB.Create(finance).Error
}

// BackfillFinanceTable mengkopi data donasi dan pengeluaran ke tabel finance
func BackfillFinanceTable() error {
	var donations []models.Donation
	var expenses []models.Expense
	var financeEntries []models.Finance

	// Ambil semua donasi
	if err := config.DB.Find(&donations).Error; err != nil {
		return err
	}

	// Buat entri keuangan untuk setiap donasi
	for _, donation := range donations {
		financeEntries = append(financeEntries, models.Finance{
			Date:            time.Now(),
			TransactionType: "donation",
			Amount:          donation.Amount, // Ambil nilai amount dari donasi
			DonationID:      &donation.ID,
			ExpenseID:       nil,
		})
	}

	// Ambil semua pengeluaran
	if err := config.DB.Find(&expenses).Error; err != nil {
		return err
	}

	// Buat entri keuangan untuk setiap pengeluaran
	for _, expense := range expenses {
		financeEntries = append(financeEntries, models.Finance{
			Date:            time.Now(),
			TransactionType: "expense",
			Amount:          expense.Amount, // Ambil nilai amount dari pengeluaran
			DonationID:      nil,
			ExpenseID:       &expense.ID,
		})
	}

	// Simpan semua entri ke tabel finance sekaligus
	if len(financeEntries) > 0 {
		if err := config.DB.Create(&financeEntries).Error; err != nil {
			return err
		}
	}

	return nil
}
