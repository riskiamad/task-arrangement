package role

import (
	"net/http"
	"strconv"
	model "task-scheduler/datamodel"
	"task-scheduler/repository"
	"task-scheduler/util"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) URLMapping(r *gin.RouterGroup) {
	r.GET("", h.read)
	r.GET("/:id", h.readDetail)
	r.POST("", h.create)
	r.PUT("/:id", h.update)

}

func (h *Handler) read(c *gin.Context) {
	pagination, err := util.GeneratePaginationFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	role := new(model.Role)
	roles, total, err := repository.GetRoles(role, pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  roles,
		"total": total,
	})

}

func (h *Handler) readDetail(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	role, err := repository.GetRole(int64(roleID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})

}

func (h *Handler) create(c *gin.Context) {
	var r createRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	newRole, err := create(r)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, newRole)
}

func (h *Handler) update(c *gin.Context) {
	var r updateRequest

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	r.ID = int64(intID)

	if err := c.ShouldBindJSON(&r); err != nil {
		println(err.Error())
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	updatedRole, err := update(r)

	if err != nil {
		errorMessage := gin.H{"error": err}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, updatedRole)
}
