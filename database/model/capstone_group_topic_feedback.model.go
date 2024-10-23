package model

import "time"

type CapstoneGroupTopicFeedback struct {
	ID                   int64              `gorm:"primaryKey;column:id;autoIncrement"`
	Feedback             string             `gorm:"column:feedback;type:text;not null"`
	ReviewerID           int64              `gorm:"column:reviewer_id;type:bigint"`
	Reviewer             Teacher            `gorm:"foreignKey:RejectedByID;references:ID"`
	CreatedAt            time.Time          `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time          `gorm:"column:updated_at;autoUpdateTime"`
	CapstoneGroupTopicID int64              `gorm:"column:capstone_group_topic_id;type:bigint"`
	CapstoneGroupTopic   CapstoneGroupTopic `gorm:"foreignKey:CapstoneGroupTopicID;references:ID"`
}

func (CapstoneGroupTopicFeedback) TableName() string {
	return "capstone_group_topic_feedbacks"
}
