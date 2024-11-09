package model

import "time"

type ScheduleReview struct {
	ID                    int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Title                 string    `gorm:"column:title;type:varchar(100);not null"`
	Description           string    `gorm:"column:description;type:text;not null"`
	LinkMeeting           string    `gorm:"column:link_meeting;type:text;not null"`
	StartTime             time.Time `gorm:"column:start_time;type:timestamp with time zone;not null"`
	EndTime               time.Time `gorm:"column:end_time;type:timestamp with time zone;not null"`
	EvaluationCommitteeID int64     `gorm:"column:evaluation_committee_id;not null"`
	EvaluationCommittee   EvaluationCommittee
	CapstoneGroupID       int64 `gorm:"column:capstone_group_id;not null"`
	CapstoneGroup         CapstoneGroup
	CreatedAt             time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt             time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ScheduleReview) TableName() string {
	return "schedule_reviews"
}
