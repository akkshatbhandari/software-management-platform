package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"software-management-platform/backend/internal/config"
	"software-management-platform/backend/internal/routes"
	"software-management-platform/backend/internal/database"
)

func main() {
	config.LoadEnv()

	database.ConnectPostgres()

	router:= gin.Default()

	routes.RegisterHealthRoutes(router)

	port := config.GetPort()

	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}