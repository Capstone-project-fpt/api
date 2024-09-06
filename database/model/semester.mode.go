package model

import (
	"time"
)

type Semester struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:text;not null"`
	StartTime time.Time `gorm:"column:start_time;type:timestamp with time zone;not null"`
	EndTime   time.Time `gorm:"column:end_time;type:timestamp with time zone;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Semester) TableName() string {
	return "semesters"
}