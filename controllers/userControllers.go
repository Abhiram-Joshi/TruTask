package controllers

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/initializers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		RoleID   int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Storing the details
	user := models.User{Email: body.Email, Password: string(hash), IsAdmin: 0, Active: 1, Deleted: 0}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Extracting user information from the database
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password@",
		})

		return
	}

	// Checking password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password!",
		})

		return
	}

	// Generating a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSecret")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func DeleteUser(c *gin.Context) {
	var body struct {
		ID uint `json:"user_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	result := initializers.DB.Delete(&models.User{}, body.ID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete User",
		})

		return
	}

	initializers.DB.Model(&models.User{}).Unscoped().Where("id = ?", body.ID).Update("deleted", 1)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Deleted",
	})
}

func DeactivateAccount(c *gin.Context) {
	var body struct {
		ID uint `json:"user_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.Unscoped().First(&user, "id = ?", body.ID)

	if user.Deleted == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User does not exist",
		})

		return
	}

	result := initializers.DB.Delete(&models.User{}, body.ID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to deactivate User",
		})

		return
	}

	initializers.DB.Model(&models.User{}).Unscoped().Where("id = ?", body.ID).Update("active", 0)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Deactivated",
	})
}

func ActivateAccount(c *gin.Context) {
	var body struct {
		ID uint `json:"user_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.Unscoped().First(&user, body.ID)

	if user.Deleted == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User does not exist",
		})

		return
	}

	initializers.DB.Model(&models.User{}).Unscoped().Where("id = ?", body.ID).Update("active", 1)
	initializers.DB.Model(&models.User{}).Unscoped().Where("id = ?", body.ID).Update("deleted_at", nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Activated",
	})
}

func AssignRoleToUser(c *gin.Context) {

	var UserAndRole struct {
		UserID uint `json:"user_id"`
		RoleID uint `json:"role_id"`
	}

	if err := c.BindJSON(&UserAndRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read body",
			"message": err,
		})

		return
	}

	fmt.Println(UserAndRole)

	var user models.User

	result_user := initializers.DB.Model(&models.User{}).First(&user, UserAndRole.UserID)

	if result_user.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Fetch User",
		})

		return
	}

	var role models.Role

	result_role := initializers.DB.Model(&models.Role{}).First(&role, UserAndRole.RoleID)

	if result_role.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Fetch Role",
		})

		return
	}

	user.Roles = append(user.Roles, &role)
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Role added successfully",
	})
}

func UploadUserFromCSV(c *gin.Context) {

	type csvFileInput struct {
		CSVFile *multipart.FileHeader `form:"file" binding:"required"`
	}

	var input csvFileInput

	if c.ShouldBind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to upload file",
		})

		return
	}

	file, err := input.CSVFile.Open()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})

		return
	}

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})

		return
	}

	// Format of CSV file
	// Email,Password,Role1 Role2 Role3

	var userRoles [][]*models.Role
	var users []models.User
	for _, record := range records {
		all_roleids := record[2]

		roleids_split := strings.Split(all_roleids, " ")

		var roles []*models.Role
		for _, role := range roleids_split {
			var temp_role *models.Role
			initializers.DB.Model(&models.Role{}).First(&temp_role, role)
			roles = append(roles, temp_role)
		}

		userRoles = append(userRoles, roles)

		hash, hash_err := bcrypt.GenerateFromPassword([]byte(record[1]), 10)

		if hash_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash a password",
			})

			return
		}

		temp_user := models.User{
			Email:    record[0],
			Password: string(hash),
		}

		users = append(users, temp_user)
	}

	initializers.DB.CreateInBatches(users, 100)

	initializers.DB.Model(&users).Association("Roles").Append(userRoles)

	defer file.Close()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Users created successfully",
	})
}
