package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/initializers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AttachUser(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWTSecret")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["userId"])

		if user.ID == 0 {
			fmt.Println("$$")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
	}

	c.Next()
}
