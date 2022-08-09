package role

import (
	"task-scheduler/config"
	model "task-scheduler/datamodel"
	"time"
)

func create(r createRequest) (m *model.Role, err error) {
	tx := config.DbConn.Begin()

	m = &model.Role{
		Name:      r.Name,
		CreatedAt: time.Now(),
	}

	if err = config.DbConn.Create(&m).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return m, nil
}

func update(r updateRequest) (m *model.Role, err error) {
	tx := config.DbConn.Begin()

	m = &model.Role{
		ID:        r.ID,
		Name:      r.Name,
		UpdatedAt: time.Now(),
	}

	if err = config.DbConn.Model(&m).Updates(&m).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return m, nil
}
