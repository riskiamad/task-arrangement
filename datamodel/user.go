package model

import "time"

type User struct {
	ID        int64     `gorm:"column(id);auto;primaryKey" json:"id"`
	Email     string    `gorm:"column(email);type:varchar(150)" json:"email"`
	Password  string    `gorm:"column(password);type:varchar(150)" json:"-"`
	CreatedAt time.Time `gorm:"column(created_at);type:timestamp;NULL;default:NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column(updated_at);type:timestamp;NULL;default:NULL" json:"updated_at"`
}
