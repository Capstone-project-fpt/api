package model

import "time"

type InvitationMentorCapstoneGroup struct {
	ID              int64         `gorm:"primaryKey;column:id;autoIncrement"`
	Status          string        `gorm:"column:status;type:varchar(50)"`
	CapstoneGroupID int64         `gorm:"column:capstone_group_id;type:bigint"`
	CapstoneGroup   CapstoneGroup `gorm:"foreignKey:CapstoneGroupID;references:ID"`
	MentorID        int64         `gorm:"column:mentor_id"`
	Mentor          Teacher       `gorm:"foreignKey:MentorID;references:ID"`
	ExpiredAt       time.Time     `gorm:"column:expired_at;type:timestamp with time zone"`
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}

func (InvitationMentorCapstoneGroup) TableName() string {
	return "invitation_mentor_capstone_groups"
}
