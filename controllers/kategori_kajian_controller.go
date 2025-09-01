package controllers

import (
	"masjid-api/models"
	"masjid-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// KategoriKajianInput adalah struct untuk input request POST/PUT
type KategoriKajianInput struct {
	Name string `json:"name" binding:"required"`
}

// GetAllKategoriKajian menangani GET /api/kategori-kajian
func GetAllKategoriKajian(c *gin.Context) {
	kategoriKajians, err := repository.GetAllKategoriKajian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve category data"})
		return
	}
	c.JSON(http.StatusOK, kategoriKajians)
}

// GetKategoriKajianByID menangani GET /api/kategori-kajian/:id
func GetKategoriKajianByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	kategoriKajian, err := repository.GetKategoriKajianByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, kategoriKajian)
}

// CreateKategoriKajian menangani POST /api/kategori-kajian
func CreateKategoriKajian(c *gin.Context) {
	var input KategoriKajianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kategoriKajian := models.KategoriKajian{Name: input.Name}
	if err := repository.CreateKategoriKajian(&kategoriKajian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, kategoriKajian)
}

// UpdateKategoriKajian menangani PUT /api/kategori-kajian/:id
func UpdateKategoriKajian(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	kategoriKajian, err := repository.GetKategoriKajianByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var input KategoriKajianInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kategoriKajian.Name = input.Name

	if err := repository.UpdateKategoriKajian(kategoriKajian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(http.StatusOK, kategoriKajian)
}

// DeleteKategoriKajian menangani DELETE /api/kategori-kajian/:id
func DeleteKategoriKajian(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := repository.DeleteKategoriKajian(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully (soft delete)"})
}
