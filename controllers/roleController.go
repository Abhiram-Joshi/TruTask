package controllers

import (
	"net/http"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/initializers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/models"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var body struct {
		Name string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errpr": "Failed to read body",
		})

		return
	}

	result := initializers.DB.Create(&models.Role{Name: body.Name})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Role",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Role Created",
	})

}

func DeleteRole(c *gin.Context) {
	var body struct {
		ID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	result := initializers.DB.Delete(&models.Role{}, body.ID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete Role",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role Deleted",
	})
}

func GetAllRoles(c *gin.Context) {

	type role_data struct {
		Name string
	}

	var roles []role_data

	result := initializers.DB.Model(&models.Role{}).Find(&roles)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Fetch Roles",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role Deleted",
		"data":    roles,
	})
}
