package controllers

import (
	"net/http"
	"strconv"
	"time"

	"masjid-api/models"
	"masjid-api/repository"
	"masjid-api/utils"

	"github.com/gin-gonic/gin"
)

// KajianInput adalah struct untuk input request POST/PUT
type KajianInput struct {
	Title            string           `json:"title" binding:"required"`
	Description      string           `json:"description"`
	Date             utils.CustomTime `json:"date" binding:"required"`
	UstadzID         uint             `json:"ustadz_id" binding:"required"`
	KategoriKajianID uint             `json:"kategori_kajian_id" binding:"required"`
}

// GetAllKajian menangani GET /api/kajian
func GetAllKajian(c *gin.Context) {
	kajians, err := repository.GetAllKajian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve kajian data"})
		return
	}
	c.JSON(http.StatusOK, kajians)
}

// GetKajianByID menangani GET /api/kajian/:id
func GetKajianByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kajian ID"})
		return
	}

	kajian, err := repository.GetKajianByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kajian not found"})
		return
	}
	c.JSON(http.StatusOK, kajian)
}

// CreateKajian menangani POST /api/kajian
func CreateKajian(c *gin.Context) {
	var input KajianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kajian := models.Kajian{
		Title:            input.Title,
		Description:      input.Description,
		Date:             time.Time(input.Date),
		UstadzID:         input.UstadzID,
		KategoriKajianID: input.KategoriKajianID,
	}

	// Langkah 1: Simpan data ke database
	if err := repository.CreateKajian(&kajian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create kajian"})
		return
	}

	// Langkah 2: Ambil kembali data yang baru disimpan dengan relasi yang lengkap
	createdKajian, err := repository.GetKajianByID(kajian.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created kajian"})
		return
	}

	c.JSON(http.StatusCreated, createdKajian)
}

// UpdateKajian menangani PUT /api/kajian/:id
func UpdateKajian(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kajian ID"})
		return
	}

	// Ambil data kajian yang ada
	kajian, err := repository.GetKajianByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kajian not found"})
		return
	}

	var input KajianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbarui field-field dari struct kajian
	kajian.Title = input.Title
	kajian.Description = input.Description
	kajian.Date = time.Time(input.Date)
	kajian.UstadzID = input.UstadzID
	kajian.KategoriKajianID = input.KategoriKajianID

	// Simpan perubahan ke database
	if err := repository.UpdateKajian(kajian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update kajian"})
		return
	}

	// Ambil kembali data yang baru diperbarui untuk mendapatkan relasi terbaru
	updatedKajian, err := repository.GetKajianByID(kajian.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated kajian"})
		return
	}

	// Berikan response dengan data yang sudah lengkap
	c.JSON(http.StatusOK, updatedKajian)
}

// DeleteKajian menangani DELETE /api/kajian/:id
func DeleteKajian(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kajian ID"})
		return
	}

	if err := repository.DeleteKajian(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete kajian"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kajian deleted successfully (soft delete)"})
}
