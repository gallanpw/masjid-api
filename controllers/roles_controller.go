package controllers

import (
	"net/http"
	"strconv"

	"masjid-api/models"
	"masjid-api/repository"

	"github.com/gin-gonic/gin"
)

type RoleInput struct {
	Name string `json:"name" binding:"required"`
}

// CreateRole handles POST /api/roles
func CreateRole(c *gin.Context) {
	var input RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := models.Role{
		Name: input.Name,
	}

	if err := repository.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// GetAllRoles handles GET /api/roles
func GetAllRoles(c *gin.Context) {
	roles, err := repository.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve roles"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// GetRoleByID handles GET /api/roles/:id
func GetRoleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := repository.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// UpdateRole handles PUT /api/roles/:id
func UpdateRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var input RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := repository.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	role.Name = input.Name

	if err := repository.UpdateRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}

	// Fetch and return the updated role to ensure consistency
	updatedRole, err := repository.GetRoleByID(role.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated role"})
		return
	}

	c.JSON(http.StatusOK, updatedRole)
}

// DeleteRole handles DELETE /api/roles/:id
func DeleteRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := repository.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
