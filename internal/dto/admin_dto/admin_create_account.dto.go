package admin_dto

type InputAdminCreateStudentAccount struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
	SubMajorID  int64  `json:"sub_major_id" binding:"required"`
}

type InputAdminCreateTeacherAccount struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	SubMajorID  int64  `json:"sub_major_id" binding:"required"`
}

type AccountWithEmail interface {
	GetEmail() string
}

func (input InputAdminCreateStudentAccount) GetEmail() string {
	return input.Email
}

func (input InputAdminCreateTeacherAccount) GetEmail() string {
	return input.Email
}