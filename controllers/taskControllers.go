package controllers

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/initializers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var body struct {
		Name        string
		Description string
		RoleID      uint `json:"role_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var role models.Role
	initializers.DB.First(&role, body.RoleID)

	task := models.Task{Name: body.Name, Description: body.Description, Done: 0, RoleID: role.ID}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create Task",
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task Created",
	})
}

func DeleteTask(c *gin.Context) {
	var body struct {
		ID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	initializers.DB.Delete(&models.Task{}, body.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Deleted",
	})
}

func GetTasks(c *gin.Context) {

	userValue, _ := c.Get("user")

	user := userValue.(models.User)

	temp_user := models.User{}

	user_result := initializers.DB.Model(&models.User{}).Preload("Roles").Find(&temp_user, user.ID)

	if user_result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Fetch Tasks",
		})

		return
	}

	var role_ids []uint

	for _, item := range temp_user.Roles {
		role_ids = append(role_ids, item.ID)
	}

	type task_schema struct {
		Name        string
		Description string
		Done        uint
	}

	task_data := []task_schema{}

	task_result := initializers.DB.Model(&models.Task{}).Find(&task_data, "role_id IN ?", role_ids)

	if task_result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Fetch Tasks",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks Fetched",
		"data":    task_data,
	})
}

func UploadTaskFromCSV(c *gin.Context) {

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
	// RoleID,Name,Description,Done

	var tasks []models.Task
	for _, record := range records {
		temp_roleid, _ := strconv.Atoi(record[0])
		temp_done, _ := strconv.Atoi(record[3])

		temp_task := models.Task{
			RoleID:      uint(temp_roleid),
			Name:        record[1],
			Description: record[2],
			Done:        uint(temp_done),
		}

		tasks = append(tasks, temp_task)
	}

	result := initializers.DB.CreateInBatches(tasks, 100)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create tasks",
		})

		return
	}

	defer file.Close()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tasks created successfully",
	})

}

func MarkTaskDone(c *gin.Context) {
	var body struct {
		ID uint `json:"task_id"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read body",
			"message": err,
		})

		return
	}

	result := initializers.DB.Model(&models.Task{}).Find(nil, body.ID).Update("done", 1)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to mark task as done",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task marked done",
	})
}
