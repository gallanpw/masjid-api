package controllers

import (
	"masjid-api/models"
	"masjid-api/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// DonationInput adalah struct untuk input request POST donasi
type DonationInput struct {
	Amount float64 `json:"amount" binding:"required,min=1"`
	Notes  string  `json:"notes"`
}

// ExpenseInput adalah struct untuk input request POST pengeluaran
type ExpenseInput struct {
	Amount float64 `json:"amount" binding:"required,min=1"`
	Notes  string  `json:"notes"`
}

// CreateDonation menangani POST /api/donations
func CreateDonation(c *gin.Context) {
	var input DonationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	donation := models.Donation{
		Amount: input.Amount,
		Notes:  input.Notes,
	}

	if err := repository.CreateDonation(&donation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create donation"})
		return
	}

	// Sekarang, buat entri keuangan dan sertakan jumlahnya
	financeTransaction := models.Finance{
		Date:            time.Now(),
		TransactionType: "donation",
		DonationID:      &donation.ID,
		Amount:          donation.Amount, // <-- **Tambahkan baris ini**
	}

	if err := repository.CreateFinanceEntry(&financeTransaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create finance record"})
		return
	}

	c.JSON(http.StatusCreated, donation)
}

// CreateExpense menangani POST /api/expenses
func CreateExpense(c *gin.Context) {
	var input ExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := models.Expense{
		Amount: input.Amount,
		Notes:  input.Notes,
	}

	if err := repository.CreateExpense(&expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	// Sekarang, buat entri keuangan dan sertakan jumlahnya
	financeTransaction := models.Finance{
		Date:            time.Now(),
		TransactionType: "expense",
		ExpenseID:       &expense.ID,
		Amount:          expense.Amount, // <-- **Tambahkan baris ini**
	}

	if err := repository.CreateFinanceEntry(&financeTransaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create finance record"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

// GetAllFinance menangani GET /api/finance
func GetAllFinance(c *gin.Context) {
	// Ambil parameter filter dari URL
	filter := c.Query("filter")
	var donations []models.Donation
	var expenses []models.Expense
	var err error

	// Tentukan rentang waktu berdasarkan filter
	now := time.Now()
	switch filter {
	case "daily":
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		endOfDay := startOfDay.Add(24 * time.Hour).Add(-1 * time.Nanosecond)
		donations, err = repository.GetAllDonationsInDateRange(startOfDay, endOfDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve donations"})
			return
		}
		expenses, err = repository.GetAllExpensesInDateRange(startOfDay, endOfDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve expenses"})
			return
		}
	case "weekly":
		startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
		endOfWeek := startOfWeek.AddDate(0, 0, 7)
		donations, err = repository.GetAllDonationsInDateRange(startOfWeek, endOfWeek)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve donations"})
			return
		}
		expenses, err = repository.GetAllExpensesInDateRange(startOfWeek, endOfWeek)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve expenses"})
			return
		}
	default: // Tanpa filter, ambil semua data
		donations, err = repository.GetAllDonations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve donations"})
			return
		}
		expenses, err = repository.GetAllExpenses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve expenses"})
			return
		}
	}

	// Hitung total saldo
	var totalDonation float64
	for _, d := range donations {
		totalDonation += d.Amount
	}

	var totalExpense float64
	for _, e := range expenses {
		totalExpense += e.Amount
	}

	totalBalance := totalDonation - totalExpense

	// Buat response gabungan
	response := gin.H{
		"total_balance": totalBalance,
		"donations":     donations,
		"expenses":      expenses,
	}

	c.JSON(http.StatusOK, response)
}
