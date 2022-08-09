package model

import "time"

type Staff struct {
	ID          int64  `gorm:"column(id);auto;primaryKey" json:"id"`
	Code        string `gorm:"column(code);type:varchar(50)" json:"code"`
	FullName    string `gorm:"column(full_name);type:varchar(100)" json:"full_name"`
	DisplayName string `gorm:"column(display_name);type:varchar(100)" json:"display_name"`
	RoleID      int64  `gorm:"column(role_id)"`
	Role        Role   `gorm:"foreignKey:RoleID;AssociationForeignKey:ID" json:"role"`
	// User        *User
	CreatedAt time.Time `gorm:"column(created_at);type:timestamp;NULL;default:NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column(updated_at);type:timestamp;NULL;default:NULL" json:"updated_at"`
}
