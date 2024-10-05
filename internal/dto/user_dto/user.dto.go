package user_dto

import "time"

type UserOutput struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
}

type StudentInfoOutput struct {
	StudentID       int       `json:"student_id"`
	Code            string    `json:"code"`
	SubMajorId      int       `json:"sub_major_id"`
	CapstoneGroupID int       `json:"capstone_group_id"`
	CreatedAt       time.Time `json:"created_at"`
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

// This used for swagger
type GetUserSwaggerOutput struct {
	Code    int            `json:"code"`
	Success bool           `json:"message"`
	Data    *GetUserOutput `json:"data"`
}
