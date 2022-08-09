package repository

import (
	"task-scheduler/config"
	model "task-scheduler/datamodel"
)

// GetRoles : function to get all role data
func GetRoles(role *model.Role, pagination *model.Pagination) (mx []*model.Role, total int64, err error) {
	m := new(model.Role)
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DbConn.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)
	result := queryBuilder.Model(&m).Where(role).Find(&mx).Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return mx, total, nil
}

// GetRole : function to get filtered role data
func GetRole(roleID int64) (role *model.Role, err error) {
	m := new(model.Role)
	result := config.DbConn.Model(&m).Where("id = ?", roleID).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}

// ValidRole : function to check if id is valid in database
func ValidRole(id int64) (m *model.Role, e error) {
	m = &model.Role{ID: id}
	e = config.DbConn.First(&m).Error

	return
}
