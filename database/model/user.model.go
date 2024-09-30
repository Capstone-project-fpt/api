package model

import (
	"time"
)

type User struct {
	ID          int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Name        string    `gorm:"column:name;type:text;not null"`
	UserType    string    `gorm:"column:user_type;type:text;not null"`
	Password    string    `gorm:"column:password;type:text"`
	Email       string    `gorm:"column:email;type:text;not null"`
	PhoneNumber string    `gorm:"column:phone_number;type:text;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Roles       []Role    `gorm:"many2many:users_roles;"`
}

func (User) TableName() string {
	return "users"
}
