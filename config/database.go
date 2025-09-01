package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		// log.Println("Error loading .env file. Assuming environment variables are set in Railway.")
	}

	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	// Prioritaskan DATABASE_URL yang disediakan oleh Railway
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Jika tidak ada, gunakan variabel lingkungan manual (untuk kompatibilitas)
		// dsn = fmt.Sprintf(
		// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		// 	os.Getenv("DB_HOST"),
		// 	os.Getenv("DB_PORT"),
		// 	os.Getenv("DB_USER"),
		// 	os.Getenv("DB_PASSWORD"),
		// 	os.Getenv("DB_NAME"),
		// 	os.Getenv("DB_SSLMODE"),
		// )
		log.Fatal("DATABASE_URL not found. Please set it in your Railway variables.")
		// log.Println("DATABASE_URL not found. Please set it in your Railway variables.")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic(err)
		log.Fatalf("Failed to connect to database: %v", err)
		// log.Println("Failed to connect to database:", err)
	}

	// log.Fatal("Database connection successfully opened.")
	log.Println("Database connection successfully opened.")
	// fmt.Println("Database connection successfully opened.")
}
