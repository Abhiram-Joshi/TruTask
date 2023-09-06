package middlewares

import (
	"net/http"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/models"
	"github.com/gin-gonic/gin"
)

func CheckAdmin(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.User).IsAdmin == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
