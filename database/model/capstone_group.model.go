package model

import (
	"time"
)

type CapstoneGroup struct {
	ID         int64     `gorm:"primaryKey;column:id;autoIncrement"`
	NameGroup  string    `gorm:"column:name_group;type:text;not null"`
	Topic      string    `gorm:"column:topic;type:text;not null"`
	MajorID    int64     `gorm:"not null"`                             // Foreign Key for Major
	Major      Major     `gorm:"foreignKey:MajorID;references:ID"`     // Many-to-One with Major
	SemesterID int64     `gorm:"not null"`                             // Foreign Key for Semester
	Semester   Semester  `gorm:"foreignKey:SemesterID;references:ID"`  // Many-to-One with Semester
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (CapstoneGroup) TableName() string {
	return "capstone_groups"
}