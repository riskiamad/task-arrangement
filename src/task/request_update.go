package task

import (
	"errors"
	model "task-scheduler/datamodel"
	"task-scheduler/repository"
	"time"
)

type updateRequest struct {
	ID          int64        `json:"-" binding:"required"`
	Description string       `json:"description" binding:"required"`
	StaffID     int64        `json:"staff_id" binding:"required"`
	StrDeadline string       `json:"deadline" binding:"required"`
	Deadline    time.Time    `json:"-"`
	Staff       *model.Staff `json:"-"`
	Task        *model.Task  `json:"-"`
}

func (c *updateRequest) Validate() error {
	var err error

	c.Task, err = repository.ValidTask(c.ID)
	if err != nil {
		return err
	}

	c.Staff, err = repository.ValidStaff(c.StaffID)
	if err != nil {
		return err
	}

	c.Deadline, err = time.Parse("2006-01-02", c.StrDeadline)
	if err != nil {
		return errors.New("invalid deadline date format")
	}

	return nil
}
