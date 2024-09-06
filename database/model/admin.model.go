package model

import (
	"time"
)

type Admin struct {
	ID              int64         `gorm:"primaryKey;autoIncrement"`
	UserID          int64         `gorm:"not null"`                              
	User            User          `gorm:"foreignKey:UserID;references:ID"`         
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}

func (Admin) TableName() string {
	return "admins"
}