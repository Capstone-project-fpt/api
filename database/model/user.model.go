package model

import (
	"time"
)

type User struct {
	ID          int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Name        string    `gorm:"column:name;type:text;not null"`
	UserType    string    `gorm:"column:user_type;type:text;not null"`
	Password    string    `gorm:"column:password;type:text"`
	Email       string    `gorm:"column:email;type:text;not null"`
	PhoneNumber string    `gorm:"column:phone_number;type:text;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Roles       []Role    `gorm:"many2many:users_roles;"`
}

type UserWithDetails struct {
	UserID          int    `gorm:"column:user_id"`
	UserName        string `gorm:"column:user_name"`
	UserEmail       string `gorm:"column:user_email"`
	UserPhoneNumber string `gorm:"column:user_phone_number"`
	UserType        string `gorm:"column:user_type"`

	TeacherID         int       `gorm:"column:teacher_id"`
	TeacherSubMajorID int       `gorm:"column:teacher_sub_major_id"`
	TeacherCreatedAt  time.Time `gorm:"column:teacher_created_at"`

	StudentID              int       `gorm:"column:student_id"`
	StudentCode            string    `gorm:"column:student_code"`
	StudentSubMajorID      int       `gorm:"column:student_sub_major_id"`
	StudentCapstoneGroupID int       `gorm:"column:student_capstone_group_id"`
	StudentCreatedAt       time.Time `gorm:"column:student_created_at"`
}

func (User) TableName() string {
	return "users"
}
