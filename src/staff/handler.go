package staff

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

	staff := new(model.Staff)
	staffs, total, err := repository.GetStaffs(staff, pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  staffs,
		"total": total,
	})

}

func (h *Handler) readDetail(c *gin.Context) {
	id := c.Param("id")
	staffID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	staff, err := repository.GetStaff(int64(staffID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": staff,
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

	newStaff, err := create(r)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, newStaff)
}

func (h *Handler) update(c *gin.Context) {
	var r updateRequest

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
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

	updatedStaff, err := update(r)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, updatedStaff)
}
