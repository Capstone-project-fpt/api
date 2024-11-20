package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Int64Array []int64

func (a Int64Array) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Int64Array) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	return nil
}

type EvaluationCommittee struct {
	ID             int64      `gorm:"primaryKey;column:id;autoIncrement"`
	Name           string     `gorm:"column:name;type:text;not null"`
	TeacherIDs     Int64Array `gorm:"column:teacher_ids;type:jsonb;default:'[]'::jsonb"`
	AssignGroupIDs Int64Array `gorm:"column:assign_group_ids;type:jsonb;default:'[]'::jsonb"`
	SemesterID     int64      `gorm:"column:semester_id;not null"`
	Semester       Semester   `gorm:"foreignKey:SemesterID;references:ID"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;autoUpdateTime"`
}

func (EvaluationCommittee) TableName() string {
	return "evaluation_committees"
}
