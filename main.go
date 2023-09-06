package main

import (
	"os"

	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/initializers"
	"github.com/Abhiram-Joshi/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	api := r.Group("api")

	routers.RegisterUserRoutes(api)
	routers.RegisterTaskRoutes(api)
	routers.RegisterRoleRoutes(api)

	port := os.Getenv("PORT")
	conn_url := "0.0.0.0:" + port
	r.Run(conn_url)
}
