package user_dto

import "time"

type InputGetUser struct {
	ID int `form:"id" binding:"required" example:"1"`
}

type OutputUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
}

type OutputStudentInfo struct {
	StudentID       int       `json:"student_id"`
	Code            string    `json:"code"`
	SubMajorId      int       `json:"sub_major_id"`
	CapstoneGroupID int       `json:"capstone_group_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type OutputTeacherInfo struct {
	TeacherID  int       `json:"teacher_id"`
	SubMajorID int       `json:"sub_major_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type OutputAdminInfo struct{}

type ExtraInfo struct {
	Student *OutputStudentInfo `json:"student,omitempty"`
	Teacher *OutputTeacherInfo `json:"teacher,omitempty"`
	Admin   *OutputAdminInfo   `json:"admin,omitempty"`
}

type OutputGetUser struct {
	CommonInfo *OutputUser `json:"common_info"`
	ExtraInfo  *ExtraInfo  `json:"extra_info"`
}

// This used for swagger
type OutputGetUserSwagger struct {
	Code    int            `json:"code"`
	Success bool           `json:"message"`
	Data    *OutputGetUser `json:"data"`
}
