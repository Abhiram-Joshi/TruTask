package routers

import (
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/controllers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.RouterGroup) {
	roleRouter := r.Group("/role")
	{
		roleRouter.POST("/create", middlewares.AttachUser, middlewares.CheckAdmin, controllers.CreateRole)
		roleRouter.DELETE("/delete", middlewares.AttachUser, middlewares.CheckAdmin, controllers.DeleteRole)
		roleRouter.GET("/get-all", middlewares.AttachUser, controllers.GetAllRoles)
	}
}
