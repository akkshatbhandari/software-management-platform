package projects

import(
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	repo := NewRepository(db)
	handler := NewHandler(repo)

	router.GET("/projects", handler.GetProjects)
	router.POST("/projects", handler.CreateProject)
}

