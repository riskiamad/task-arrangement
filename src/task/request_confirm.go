package task

import (
	"errors"
	model "task-scheduler/datamodel"
	"task-scheduler/repository"
	"time"
)

type confirmRequest struct {
	ID         int64       `json:"-" binding:"required"`
	Task       *model.Task `json:"-"`
	FinishedAt time.Time   `json:"-"`
}

func (c *confirmRequest) Validate() error {
	var err error

	c.Task, err = repository.ValidTask(c.ID)
	if err != nil {
		return err
	}

	if c.Task.IsDone == 1 {
		return errors.New("task already confirmed")
	}

	currentTime := time.Now()
	timeStr := currentTime.Format("2006-01-02")

	c.FinishedAt, err = time.Parse("2006-01-02", timeStr)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
