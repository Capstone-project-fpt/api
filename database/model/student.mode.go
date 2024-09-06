package model

import (
	"time"
)

type Student struct {
	ID              int64         `gorm:"primaryKey;autoIncrement"`
	Code            string        `gorm:"column:code;type:text;not null"`
	SubMajorID      int64         `gorm:"not null"`                                 // Foreign Key for SubMajor
	SubMajor        SubMajor      `gorm:"foreignKey:SubMajorID;references:ID"`      // Many-to-One with SubMajor
	UserID          int64         `gorm:"not null"`                                 // Foreign Key for User
	User            User          `gorm:"foreignKey:UserID;references:ID"`          // Many-to-One with User
	CapstoneGroupID int64         `gorm:"default:null"`                             // Foreign Key for CapstoneGroup (nullable)
	CapstoneGroup   CapstoneGroup `gorm:"foreignKey:CapstoneGroupID;references:ID"` // Many-to-One with CapstoneGroup (nullable)
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}

func (Student) TableName() string {
	return "students"
}