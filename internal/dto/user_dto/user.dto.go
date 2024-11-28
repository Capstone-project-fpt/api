package user_dto

import (
	"time"

	"github.com/api/database/model"
)

type UserOutput struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
}

func ToUserOutput(user *model.User) *UserOutput {
	return &UserOutput{
		ID:          int(user.ID),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UserType:    user.UserType,
	}
}

type StudentInfoOutput struct {
	StudentID  int       `json:"student_id"`
	Code       string    `json:"code"`
	SubMajorId int       `json:"sub_major_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type TeacherInfoOutput struct {
	TeacherID  int       `json:"teacher_id"`
	SubMajorID int       `json:"sub_major_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type AdminInfoOutput struct{}

type ExtraInfo struct {
	Student *StudentInfoOutput `json:"student,omitempty"`
	Teacher *TeacherInfoOutput `json:"teacher,omitempty"`
	Admin   *AdminInfoOutput   `json:"admin,omitempty"`
}

type GetUserOutput struct {
	CommonInfo *UserOutput `json:"common_info"`
	ExtraInfo  *ExtraInfo  `json:"extra_info"`
}

type TeacherOutput struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
	SubMajorID  int    `json:"sub_major_id"`
}

func ToTeacherOutput(teacher *model.Teacher) *TeacherOutput {
	return &TeacherOutput{
		ID:          int(teacher.ID),
		UserID:      int(teacher.UserID),
		Name:        teacher.User.Name,
		Email:       teacher.User.Email,
		PhoneNumber: teacher.User.PhoneNumber,
		UserType:    teacher.User.UserType,
		SubMajorID:  int(teacher.SubMajorID),
	}
}

type StudentOutput struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
	SubMajorID  int    `json:"sub_major_id"`
	Code        string `json:"code"`
}

func ToStudentOutput(student *model.Student) *StudentOutput {
	return &StudentOutput{
		ID:          int(student.ID),
		UserID:      int(student.UserID),
		Name:        student.User.Name,
		Email:       student.User.Email,
		PhoneNumber: student.User.PhoneNumber,
		UserType:    student.User.UserType,
		SubMajorID:  int(student.SubMajorID),
		Code:        student.Code,
	}
}

// This used for swagger
type GetUserSwaggerOutput struct {
	Code    int            `json:"code"`
	Success bool           `json:"message"`
	Data    *GetUserOutput `json:"data"`
}
