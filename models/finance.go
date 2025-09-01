package models

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	gorm.Model
	Date            time.Time `gorm:"not null"`
	Amount          float64   `json:"amount" gorm:"not null"`
	TransactionType string    `gorm:"type:varchar(50);not null"`
	DonationID      *uint     `gorm:"index"`
	ExpenseID       *uint     `gorm:"index"`
	Donation        Donation
	Expense         Expense
}
