package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   int64  `gorm:"column:id; type:int; not null; unique; primaryKey; autoIncrement;"`
	Name string `gorm:"column:name; not null, type:varchar(255)"`
	Note string `gorm:"column:note; type:varchar(255)"`
}

func (u *Role) TableName() string {
	return "roles"
}
