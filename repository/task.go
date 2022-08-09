package repository

import (
	"task-scheduler/config"
	model "task-scheduler/datamodel"
)

// GetTasks : function to get all task data
func GetTasks(task *model.Task, pagination *model.Pagination) (mx []*model.Task, total int64, err error) {
	m := new(model.Task)
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DbConn.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)
	result := queryBuilder.Model(&m).Where(task).Preload("Assignee").Find(&mx).Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return mx, total, nil
}

// GetTask : function to get filtered task data
func GetTask(taskID int64) (task *model.Task, err error) {
	m := new(model.Task)
	result := config.DbConn.Model(&m).Where("id = ?", taskID).Find(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

// ValidTask : function to check if id is valid in database
func ValidTask(id int64) (m *model.Task, e error) {
	m = &model.Task{ID: id}
	e = config.DbConn.First(&m).Error

	return
}
