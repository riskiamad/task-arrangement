package task

import (
	"task-scheduler/config"
	model "task-scheduler/datamodel"
	"time"
)

func create(r createRequest) (m *model.Task, err error) {
	tx := config.DbConn.Begin()

	m = &model.Task{
		Description: r.Description,
		StaffID:     r.StaffID,
		Deadline:    r.Deadline,
		IsDone:      2,
		CreatedAt:   time.Now(),
	}

	if err = config.DbConn.Create(&m).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return m, nil
}

func update(r updateRequest) (m *model.Task, err error) {
	tx := config.DbConn.Begin()

	r.Task.Description = r.Description
	r.Task.StaffID = r.StaffID
	r.Task.Deadline = r.Deadline
	r.Task.UpdatedAt = time.Now()

	if err = config.DbConn.Model(&r.Task).Updates(&r.Task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return r.Task, nil
}

func confirm(r confirmRequest) (m *model.Task, err error) {
	tx := config.DbConn.Begin()

	r.Task.FinishedAt = r.FinishedAt
	r.Task.IsDone = 1

	if err = config.DbConn.Model(&r.Task).Updates(&r.Task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return r.Task, nil
}
