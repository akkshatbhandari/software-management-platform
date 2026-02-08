package main

import (
	"log"

	"software-management-platform/services/core-go/internal/auth"
	"software-management-platform/services/core-go/internal/config"
	"software-management-platform/services/core-go/internal/database"
	"software-management-platform/services/core-go/internal/projects"
	"software-management-platform/services/core-go/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectPostgres()

	router:= gin.Default()

	projects.RegisterRoutes(router, database.DB)

	authRepo := &auth.Repository{DB: database.DB}

	auth.RegisterAuthRoutes(router, authRepo)

	routes.RegisterHealthRoutes(router)

	port := config.GetPort()

	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}