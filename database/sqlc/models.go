// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"
)

type CapstoneGroup struct {
	ID         int64
	NameGroup  string
	Topic      string
	MajorID    int64
	SemesterID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Comment struct {
	ID         int64
	Message    string
	DocumentID int64
	GroupID    int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Document struct {
	ID              int64
	Name            string
	FileIds         []string
	CapstoneGroupID int64
	Score           sql.NullInt16
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type File struct {
	ID        int64
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Major struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permission struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolesPermission struct {
	ID           int64
	RoleID       int64
	PermissionID int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Semester struct {
	ID        int64
	Name      string
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Student struct {
	ID              int64
	Code            string
	SubMajorID      int64
	UserID          int64
	CapstoneGroupID sql.NullInt64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type SubMajor struct {
	ID        int64
	Name      string
	MajorID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          int64
	Name        string
	UserType    string
	Password    sql.NullString
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UsersRole struct {
	ID        int64
	RoleID    int64
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
