package handler

import (
	"net/http"
	"task-scheduler/config"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct{}

func (h *RoleHandler) URLMapping(r *gin.RouterGroup) {
	r.GET("", h.index)
	r.GET("/:id", h.detail)
}

func (h *RoleHandler) index(c *gin.Context) {
	roles, err := http.Get(config.Config.APIBaseURL + "/role")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	defer roles.Body.Close()

	c.HTML(http.StatusOK, "role_index.html", gin.H{"roles": roles})
}

func (h *RoleHandler) detail(c *gin.Context) {
	id := c.Param("id")

	role, err := http.Get(config.Config.APIBaseURL + "/role/" + id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	defer role.Body.Close()

	c.HTML(http.StatusOK, "role_index.html", gin.H{"role": role})
}
