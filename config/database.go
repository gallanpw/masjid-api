package config

import (
	"fmt"
	"log"
	"masjid-api/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah variabel koneksi database global.
var DB *gorm.DB

func ConnectDB() {
	// log.Println("Loading .env file...")
	// Mencoba memuat .env file (hanya untuk pengembangan lokal)
	// err := godotenv.Load()
	_ = godotenv.Load()
	log.Println("Environment variables loaded (or skipped).")

	// if err != nil {
	// 	// log.Fatalf("Error loading .env file")
	// 	log.Println("Error loading .env file. Assuming environment variables are set in Railway.")
	// }

	// Prioritaskan DATABASE_URL yang disediakan oleh Railway
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Jika tidak ada, gunakan variabel lingkungan manual (untuk kompatibilitas)
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		)
	}

	// Menggunakan perulangan retry untuk menghubungkan ke database
	maxRetries := 5
	for i := 1; i <= maxRetries; i++ {
		log.Printf("Mencoba menghubungkan ke database... (Percobaan %d dari %d)", i, maxRetries)
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Gagal terhubung: %v. Mencoba lagi dalam 5 detik...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("Koneksi database berhasil. Menjalankan migrasi...")
		err = DB.AutoMigrate(
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
			log.Fatalf("Gagal melakukan migrasi tabel: %v", err)
		}

		log.Println("Migrasi database selesai dengan sukses.")
		return // Keluar dari fungsi jika berhasil
	}

	log.Fatal("Gagal menghubungkan ke database setelah beberapa kali mencoba. Keluar.")
}
