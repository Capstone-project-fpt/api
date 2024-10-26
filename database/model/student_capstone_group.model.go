package model

import "time"

type StudentCapstoneGroup struct {
	ID              int64         `gorm:"primaryKey;autoIncrement"`
	StudentID       int64         `gorm:"column:student_id;not null"`
	Student         Student       `gorm:"foreignKey:StudentID;references:ID"`
	CapstoneGroupID int64         `gorm:"column:capstone_group_id;not null"`
	CapstoneGroup   CapstoneGroup `gorm:"foreignKey:CapstoneGroupID;references:ID"`
	SemesterID      int64         `gorm:"column:semester_id;not null"`
	Semester        Semester      `gorm:"foreignKey:SemesterID;references:ID"`
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}

func (StudentCapstoneGroup) TableName() string {
	return "student_capstone_groups"
}