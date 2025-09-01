package main

import (
	// "fmt"
	// "log"
	"masjid-api/config"
	// "masjid-api/repository"
	"masjid-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// migrasi database secara otomatis
	config.AutoMigrateTables()

	// --- PANGGIL FUNGSI BACKFILL DI SINI ---
	// Hapus baris ini setelah pertama kali dijalankan
	// if err := repository.BackfillFinanceTable(); err != nil {
	// 	log.Fatalf("Failed to backfill finance table: %v", err)
	// }
	// fmt.Println("Finance table backfilled successfully.")
	// --- AKHIR PANGGILAN FUNGSI BACKFILL ---

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
	r.Run(":" + port)
}
