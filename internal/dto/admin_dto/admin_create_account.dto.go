package admin_dto

type InputAdminCreateStudentAccount struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
	SubMajorId  int64  `json:"sub_major_id" binding:"required"`
}
