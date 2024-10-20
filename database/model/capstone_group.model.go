package model

import (
	"time"
)

type CapstoneGroup struct {
	ID         int64     `gorm:"primaryKey;column:id;autoIncrement"`
	NameGroup  string    `gorm:"column:name_group;type:text;not null"`
	Topic      string    `gorm:"column:topic;type:text;"`
	MajorID    int64     `gorm:"not null"`
	Major      Major     `gorm:"foreignKey:MajorID;references:ID"`
	SemesterID int64     `gorm:"not null"`
	Semester   Semester  `gorm:"foreignKey:SemesterID;references:ID"`
	LeaderID   int64     `gorm:"column:leader_id;not null"`
	MentorID   int64     `gorm:"column:mentor_id"`
	Mentor     Teacher   `gorm:"foreignKey:MentorID;references:ID"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (CapstoneGroup) TableName() string {
	return "capstone_groups"
}
