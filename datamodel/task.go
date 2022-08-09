package model

import "time"

type Task struct {
	ID          int64     `gorm:"column(id);auto;primaryKey" json:"id"`
	Description string    `gorm:"column(description);type:varchar(250)" json:"description"`
	StaffID     int64     `gorm:"column(assignee)"`
	Assignee    *Staff    `gorm:"foreignKey:StaffID;AssociationForeignKey:ID" json:"assignee"`
	IsDone      int8      `gorm:"column(is_done);type:tinyint(2)" json:"is_done"`
	Deadline    time.Time `gorm:"column(deadline);type:timestamp;NULL;default:NULL" json:"deadline"`
	FinishedAt  time.Time `gorm:"column(finished_at);type:timestamp;NULL;default:NULL" json:"finished_at"`
	CreatedAt   time.Time `gorm:"column(created_at);type:timestamp;NULL;default:NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column(updated_at);type:timestamp;NULL;default:NULL" json:"updated_at"`
}
