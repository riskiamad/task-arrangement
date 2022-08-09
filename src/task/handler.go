package task

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
	r.PUT("/confirm/:id", h.confirm)
}

func (h *Handler) read(c *gin.Context) {
	pagination, err := util.GeneratePaginationFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := new(model.Task)
	tasks, total, err := repository.GetTasks(task, pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  tasks,
		"total": total,
	})

}

func (h *Handler) readDetail(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := repository.GetStaff(int64(taskID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
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

	err := r.Validate()
	if err != nil {
		println(err)
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	newTask, err := create(r)

	if err != nil {
		errorMessage := gin.H{"error": err}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, newTask)
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
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	err = r.Validate()
	if err != nil {
		println(err)
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	updatedTask, err := update(r)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (h *Handler) confirm(c *gin.Context) {
	var r confirmRequest

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	r.ID = int64(intID)

	if err := c.ShouldBindUri(&r); err != nil {
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	err = r.Validate()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	confirmedTask, err := confirm(r)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, confirmedTask)
}
