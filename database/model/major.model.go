package model

import (
	"time"
)

type Major struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string    `gorm:"column:name;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	SubMajors []SubMajor
}

func (Major) TableName() string {
	return "majors"
}
