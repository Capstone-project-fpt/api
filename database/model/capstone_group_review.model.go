package model

import (
	"github.com/lib/pq"
	"time"
)

type CapstoneGroupReview struct {
	ID               int64          `gorm:"primaryKey;column:id;autoIncrement"`
	CapstoneGroupID  int64          `gorm:"column:capstone_group_id;type:bigint"`
	CapstoneGroup    CapstoneGroup  `gorm:"foreignKey:CapstoneGroupID;references:ID"`
	ScheduleReviewID int64          `gorm:"column:schedule_review_id;type:bigint"`
	ScheduleReview   ScheduleReview `gorm:"foreignKey:ScheduleReviewID;references:ID"`
	ReportFiles      pq.StringArray `gorm:"column:report_files;type:text[]"`
	Feedback         string         `gorm:"column:feedback;type:text"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (CapstoneGroupReview) TableName() string {
	return "capstone_group_reviews"
}
