package model

import "time"

type Role struct {
	ID        int64     `gorm:"column(id);auto;primaryKey" json:"id"`
	Name      string    `gorm:"column(name);type:varchar(100)" json:"name"`
	CreatedAt time.Time `gorm:"column(created_at);type:timestamp;NULL;default:NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column(updated_at);type:timestamp;NULL;default:NULL" json:"updated_at"`
}
