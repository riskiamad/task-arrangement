package staff

import (
	model "task-scheduler/datamodel"
	"task-scheduler/repository"
)

type updateRequest struct {
	ID          int64       `json:"-" binding:"required"`
	FullName    string      `json:"full_name" binding:"required"`
	DisplayName string      `json:"display_name" binding:"required"`
	RoleID      int64       `json:"role_id" binding:"required"`
	Role        *model.Role `json:"-"`
}

func (c *updateRequest) Validate() error {
	var err error

	c.Role, err = repository.ValidRole(c.RoleID)
	if err != nil {
		return err
	}

	return nil
}
