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

func (h *Handler) CreateProject(c *gin.Context) {
	var input CreateProjectInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "invalid request body",
		})
		return 
	}

	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "project name is required",
		})
		return
	}

	project, err := h.Repo.Create(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create project",
		})
		return
	}

	c.JSON(http.StatusCreated, project)
}