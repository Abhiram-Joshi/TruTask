package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/initializers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/models"
	"github.com/gin-gonic/gin"
)

func CheckSameRole(c *gin.Context) {
	var body struct {
		RoleID uint `json:"role_id"`
	}

	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user, _ := c.Get("user")
	userInstance := user.(models.User)

	var Userroles []models.Role
	initializers.DB.Model(&models.User{}).Preload("Roles").Find(&Userroles, userInstance.ID)

	fmt.Println(Userroles)
	for _, role := range Userroles {
		fmt.Println(role.ID)
		if role.ID == body.RoleID {
			c.Next()
		}
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}
