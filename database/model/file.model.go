package model

import (
	"time"
)

type File struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Path      string    `gorm:"column:path;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (File) TableName() string {
	return "files"
}
