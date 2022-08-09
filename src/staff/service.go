package staff

import (
	"task-scheduler/config"
	model "task-scheduler/datamodel"
	"task-scheduler/util"
	"time"
)

func create(r createRequest) (m *model.Staff, err error) {
	tx := config.DbConn.Begin()

	r.Code, err = util.GenerateCode("STF", "staff")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	m = &model.Staff{
		Code:        r.Code,
		FullName:    r.FullName,
		DisplayName: r.DisplayName,
		RoleID:      r.RoleID,
		CreatedAt:   time.Now(),
	}

	if err = config.DbConn.Create(&m).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return m, nil
}

func update(r updateRequest) (m *model.Staff, err error) {
	tx := config.DbConn.Begin()

	m = &model.Staff{
		ID:          r.ID,
		FullName:    r.FullName,
		DisplayName: r.DisplayName,
		RoleID:      r.RoleID,
		UpdatedAt:   time.Now(),
	}

	if err = config.DbConn.Model(&m).Updates(&m).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return m, nil
}
