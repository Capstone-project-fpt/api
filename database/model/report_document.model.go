package model

import (
	"time"

	"github.com/lib/pq"
)

type ReportDocument struct {
	ID                 int64          `gorm:"primaryKey;column:id;autoIncrement"`
	Name               string         `gorm:"column:name;type:text;not null"`
	FileIDs            pq.StringArray `gorm:"column:file_ids;type:text[]"`
	CapstoneGroupID    int64          `gorm:"column:capstone_group_id;type:bigint"`
	CapstoneGroup      CapstoneGroup  `gorm:"foreignKey:CapstoneGroupID;references:ID"`
	MentorReviewStatus string         `gorm:"column:mentor_review_status;type:varchar(50)"`
	TypeReport         string         `gorm:"column:type_report;type:varchar(50)"`
	Conclusion         *string        `gorm:"column:conclusion;type:text"`
	CreatedAt          time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (ReportDocument) TableName() string {
	return "report_documents"
}
