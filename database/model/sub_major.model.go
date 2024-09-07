package model

import (
	"time"
)

type SubMajor struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string    `gorm:"column:name;type:text;not null"`
	MajorID   int64     `gorm:"column:major_id;not null"`
	Major     Major     `gorm:"foreignKey:MajorID;references:ID"` // Establishes the many-to-one relationship
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (SubMajor) TableName() string {
	return "sub_majors"
}