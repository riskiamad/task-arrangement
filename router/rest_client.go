package router

import (
	"task-scheduler/web/handler"
)

func init() {
	handlers["staff"] = &handler.StaffHandler{}
	handlers["role"] = &handler.RoleHandler{}
	handlers["task"] = &handler.TaskHandler{}
}
