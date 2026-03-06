package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(router *gin.Engine, repo *Repository) {
	auth := router.Group("/auth")

	auth.POST("/register", Register(repo))
	auth.POST("/login", Login(repo))
	auth.POST("/refresh/store", StoreRefreshToken(repo))
	auth.POST("/refresh/validate", ValidateRefreshToken(repo))
	auth.POST("/refresh/revoke", RevokeRefreshToken(repo))
}
