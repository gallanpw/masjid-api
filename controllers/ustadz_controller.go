package controllers

import (
	"masjid-api/models"
	"masjid-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UstadzInput adalah struct untuk input request POST/PUT
type UstadzInput struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio"`
}

// GetAllUstadz menangani GET /api/ustadz
func GetAllUstadz(c *gin.Context) {
	ustadzs, err := repository.GetAllUstadz()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ustadz data"})
		return
	}
	c.JSON(http.StatusOK, ustadzs)
}

// GetUstadzByID menangani GET /api/ustadz/:id
func GetUstadzByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ustadz ID"})
		return
	}

	ustadz, err := repository.GetUstadzByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ustadz not found"})
		return
	}
	c.JSON(http.StatusOK, ustadz)
}

// CreateUstadz menangani POST /api/ustadz
func CreateUstadz(c *gin.Context) {
	var input UstadzInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ustadz := models.Ustadz{
		Name: input.Name,
		Bio:  input.Bio,
	}

	if err := repository.CreateUstadz(&ustadz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ustadz"})
		return
	}
	c.JSON(http.StatusCreated, ustadz)
}

// UpdateUstadz menangani PUT /api/ustadz/:id
func UpdateUstadz(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ustadz ID"})
		return
	}

	ustadz, err := repository.GetUstadzByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ustadz not found"})
		return
	}

	var input UstadzInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ustadz.Name = input.Name
	ustadz.Bio = input.Bio

	if err := repository.UpdateUstadz(ustadz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ustadz"})
		return
	}
	c.JSON(http.StatusOK, ustadz)
}

// DeleteUstadz menangani DELETE /api/ustadz/:id
func DeleteUstadz(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ustadz ID"})
		return
	}

	if err := repository.DeleteUstadz(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ustadz"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ustadz deleted successfully (soft delete)"})
}
