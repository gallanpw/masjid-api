package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// Mencoba memuat .env file (hanya untuk pengembangan lokal)
	// err := godotenv.Load()
	_ = godotenv.Load()

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

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic(err)
		// log.Fatalf("Failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Database connection successfully opened.")
	return nil
}
