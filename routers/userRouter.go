package routers

import (
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/controllers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {

	userRouter := r.Group("/user")
	{
		userRouter.POST("/signup", controllers.Signup)
		userRouter.POST("/login", controllers.Login)
		userRouter.DELETE("/delete", middlewares.AttachUser, middlewares.CheckAdmin, controllers.DeleteUser)
		userRouter.POST("/assign-role", middlewares.AttachUser, middlewares.CheckAdmin, controllers.AssignRoleToUser)
		userRouter.POST("/upload-file", middlewares.AttachUser, middlewares.CheckAdmin, controllers.UploadUserFromCSV)
		userRouter.POST("/deactivate", middlewares.AttachUser, middlewares.CheckAdmin, controllers.DeactivateAccount)
		userRouter.POST("/activate", middlewares.AttachUser, middlewares.CheckAdmin, controllers.ActivateAccount)
	}
}
