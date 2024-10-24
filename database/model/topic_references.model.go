package model

import "time"

type TopicReferences struct {
	ID                 int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Name               string    `gorm:"column:name;type:text;not null"`
	StatusReview       string    `gorm:"column:status_review;type:text;"`
	Path               string    `gorm:"column:path;type:text;not null"`
	TeacherID          int64     `gorm:"column:teacher_id;not null"`
	Teacher            Teacher   `gorm:"foreignKey:TeacherID;references:ID"`
	TopicReferencesTvs string    `gorm:"column:topic_references_tvs;tsvector"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (TopicReferences) TableName() string {
	return "topic_references"
}
