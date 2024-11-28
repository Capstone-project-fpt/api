package model

import "time"

type VerifierTopic struct {
	ID         int64     `gorm:"primaryKey;column:id;autoIncrement"`
	TeacherID  int64     `gorm:"column:teacher_id;not null"`
	Teacher    Teacher   `gorm:"foreignKey:TeacherID;references:ID"`
	SemesterID int64     `gorm:"not null"`
	Semester   Semester  `gorm:"foreignKey:SemesterID;references:ID"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (VerifierTopic) TableName() string {
	return "verifier_topics"
}
