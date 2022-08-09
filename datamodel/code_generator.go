package model

type CodeGenerator struct {
	ID    int64  `gorm:"column(id);auto;primaryKey" json:"id"`
	Table string `gorm:"column(table);type:varchar(100)" json:"table"`
	Value string `gorm:"column(value);type:varchar(100)" json:"value"`
}
