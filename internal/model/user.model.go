package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64 `gorm:"column:id; type:int; not null; unique; primaryKey; autoIncrement;"`
	UserName string    `gorm:"column:user_name;"`
	IsActive bool      `gorm:"column:is_active;"`
	Roles    []Role    `gorm:"many2many:users_roles;"`
}

func (u *User) TableName() string {
	return "users"
}
