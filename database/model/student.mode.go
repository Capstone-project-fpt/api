package model

import (
	"time"
)

type Student struct {
	ID              int64         `gorm:"primaryKey;autoIncrement"`
	Code            string        `gorm:"column:code;type:text;not null"`
	SubMajorID      int64         `gorm:"not null"`                                 		
	SubMajor        SubMajor      `gorm:"foreignKey:SubMajorID;references:ID"`      		
	UserID          int64         `gorm:"not null"`                                 		
	User            User          `gorm:"foreignKey:UserID;references:ID"`          		
	CapstoneGroupID int64         `gorm:"default:null"`                             		
	CapstoneGroup   CapstoneGroup `gorm:"foreignKey:CapstoneGroupID;references:ID"` 		
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}

func (Student) TableName() string {
	return "students"
}