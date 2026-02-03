package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) GetProjects(c *gin.Context) {
	projects, err := h.Repo.GetAll()
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve projects",
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}