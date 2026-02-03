package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"software-management-platform/services/core-go/internal/config"
	"software-management-platform/services/core-go/internal/routes"
	"software-management-platform/services/core-go/internal/database"
	"software-management-platform/services/core-go/internal/projects"
)

func main() {
	config.LoadEnv()

	database.ConnectPostgres()

	router:= gin.Default()

	projects.RegisterRoutes(router, database.DB)

	routes.RegisterHealthRoutes(router)

	port := config.GetPort()

	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}