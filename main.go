package main

import (
	"log"
	"masjid-api/config"

	// "masjid-api/repository"
	"masjid-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	// config.ConnectDB()
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// migrasi database secara otomatis
	// config.AutoMigrateTables()

	// --- PANGGIL FUNGSI BACKFILL DI SINI ---
	// Hapus baris ini setelah pertama kali dijalankan
	// if err := repository.BackfillFinanceTable(); err != nil {
	// 	log.Fatalf("Failed to backfill finance table: %v", err)
	// }
	// fmt.Println("Finance table backfilled successfully.")
	// --- AKHIR PANGGILAN FUNGSI BACKFILL ---

	// Retry logic untuk koneksi database
	// maxAttempts := 5
	// for i := 1; i <= maxAttempts; i++ {
	// 	log.Printf("Attempting to connect to database... (Attempt %d of %d)", i, maxAttempts)

	// 	// Panggilan fungsi sekarang cocok dengan tanda tangan yang mengembalikan 'error'
	// 	if err := config.ConnectDB(); err != nil {
	// 		log.Printf("Failed to connect: %v. Retrying in 5 seconds...", err)
	// 		time.Sleep(5 * time.Second)
	// 	} else {
	// 		log.Println("Database connection successful.")
	// 		break
	// 	}

	// 	if i == maxAttempts {
	// 		log.Fatal("Failed to connect to database after multiple retries. Exiting.")
	// 	}
	// }

	// Inisialisasi router Gin
	r := gin.Default()

	// Setup rute API
	routes.SetupRoutes(r)

	// Menjalankan server
	// r.Run(":8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
