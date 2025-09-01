package config

import (
	"log"
	"masjid-api/models"
)

// AutoMigrateTables bertanggung jawab untuk membuat atau memperbarui semua tabel di database
func AutoMigrateTables() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Ustadz{},
		&models.KategoriKajian{},
		&models.Kajian{},
		&models.Donation{},
		&models.Expense{},
		&models.Finance{},
	)

	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}

	log.Println("Migrasi database berhasil.")
}
