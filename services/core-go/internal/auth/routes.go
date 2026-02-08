package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(router *gin.Engine, repo *Repository) {
	auth := router.Group("/auth")

	auth.POST("/register", Register(repo))
	auth.POST("/login", Login(repo))
}