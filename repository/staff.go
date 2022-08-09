package repository

import (
	"task-scheduler/config"
	model "task-scheduler/datamodel"
)

// GetStaffs : function to get all staff data
func GetStaffs(staff *model.Staff, pagination *model.Pagination) (mx []*model.Staff, total int64, err error) {
	m := new(model.Staff)
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DbConn.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)
	result := queryBuilder.Model(&m).Where(staff).Preload("Role").Find(&mx).Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return mx, total, nil
}

// GetStaff : function to get filtered staff data
func GetStaff(staffID int64) (staff *model.Staff, err error) {
	m := new(model.Staff)
	result := config.DbConn.Model(&m).Where("id = ?", staffID).Preload("Role").Find(&staff)
	if result.Error != nil {
		return nil, result.Error
	}

	return staff, nil
}

// ValidStaff : function to check if id is valid in database
func ValidStaff(id int64) (m *model.Staff, e error) {
	m = &model.Staff{ID: id}
	e = config.DbConn.First(&m).Error

	return
}
