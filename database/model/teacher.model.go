package model

import "time"

type Teacher struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	UserID     int64     `gorm:"column:user_id;not null"`
	User       User      `gorm:"foreignKey:UserID;references:ID"`
	SubMajorID int64     `gorm:"column:sub_major_id;not null"`
	SubMajor   SubMajor  `gorm:"foreignKey:SubMajorID;references:ID"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Teacher) TableName() string {
	return "teachers"
}
