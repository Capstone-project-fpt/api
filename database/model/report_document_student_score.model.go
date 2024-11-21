package model

import "time"

type ReportDocumentStudentScore struct {
	ID               int64          `gorm:"primaryKey;column:id;autoIncrement"`
	Score            *float64       `gorm:"column:score"`
	StudentID        int64          `gorm:"column:student_id;not null"`
	Student          Student        `gorm:"foreignKey:StudentID;references:ID"`
	ReportDocumentID int64          `gorm:"column:report_document_id;not null"`
	ReportDocument   ReportDocument `gorm:"foreignKey:ReportDocumentID;references:ID"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (ReportDocumentStudentScore) TableName() string {
	return "report_document_student_scores"
}
