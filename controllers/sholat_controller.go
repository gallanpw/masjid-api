package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetJadwalSholat mengambil jadwal sholat dari API shalat-api.vercel.app
// dan meneruskannya langsung sebagai respons JSON.
func GetJadwalSholat(c *gin.Context) {
	// URL API tanpa parameter kota
	url := "https://shalat-api.vercel.app/shalat"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke API eksternal"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Gagal mendapatkan data dari API eksternal"})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons API"})
		return
	}

	// Langsung kembalikan data apa adanya tanpa parsing ke struct
	c.Data(http.StatusOK, "application/json", body)
}
