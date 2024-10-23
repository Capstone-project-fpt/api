package model

import "time"

type CapstoneGroupTopic struct {
	ID              int64         `gorm:"primaryKey;column:id;autoIncrement"`
	Topic           string        `gorm:"column:topic;type:varchar(255)"`
	DocumentPath    string        `gorm:"column:document_path;type:varchar(255)"`
	StatusReview    string        `gorm:"column:status_review;type:varchar(50)"`
	ApprovedAt      time.Time     `gorm:"column:approved_at;type:timestamp with time zone"`
	ApprovedByID    int64         `gorm:"column:approved_by_id;type:bigint"`
	ApprovedBy      Teacher       `gorm:"foreignKey:ApprovedByID;references:ID"`
	RejectedAt      time.Time     `gorm:"column:rejected_at;type:timestamp with time zone"`
	RejectedByID    int64         `gorm:"column:rejected_by_id;type:bigint"`
	RejectedBy      Teacher       `gorm:"foreignKey:RejectedByID;references:ID"`
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
	CapstoneGroupID int64         `gorm:"column:capstone_group_id;type:bigint"`
	CapstoneGroup   CapstoneGroup `gorm:"foreignKey:CapstoneGroupID;references:ID"`
}

func (CapstoneGroupTopic) TableName() string {
	return "capstone_group_topics"
}
