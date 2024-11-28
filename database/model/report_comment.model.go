package model

import "time"

type ReportComment struct {
	ID               int64          `gorm:"primaryKey;column:id;autoIncrement"`
	UserID           int64          `gorm:"column:user_id;not null"`
	User             User           `gorm:"foreignKey:UserID;references:ID"`
	ReportDocumentID int64          `gorm:"column:report_document_id;not null"`
	ReportDocument   ReportDocument `gorm:"foreignKey:ReportDocumentID;references:ID"`
	Message          string         `gorm:"column:message;type:text;not null"`
	GroupComment     int64          `gorm:"column:group_comment;type:integer;not null"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (ReportComment) TableName() string {
	return "report_comments"
}
