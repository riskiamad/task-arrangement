package router

import (
	"task-scheduler/src/role"
	"task-scheduler/src/staff"
	"task-scheduler/src/task"
)

func init() {
	handlers["api/staff"] = &staff.Handler{}
	handlers["api/role"] = &role.Handler{}
	handlers["api/task"] = &task.Handler{}
}
