package routers

import (
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/controllers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.RouterGroup) {
	taskRouter := r.Group("/task")
	{
		taskRouter.POST("/create", middlewares.AttachUser, middlewares.CheckAdmin, controllers.CreateTask)
		taskRouter.DELETE("/delete", middlewares.AttachUser, middlewares.CheckAdmin, controllers.DeleteTask)
		taskRouter.GET("/get", middlewares.AttachUser, controllers.GetTasks)
		taskRouter.POST("/upload-file", middlewares.AttachUser, middlewares.CheckAdmin, controllers.UploadTaskFromCSV)
		taskRouter.POST("/mark-done", middlewares.AttachUser, controllers.MarkTaskDone)
	}
}
