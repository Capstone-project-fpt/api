package admin_dto

import (
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/user_dto"
)

type GetListUsersInput struct {
	Limit     int      `form:"limit" binding:"required" example:"10"`
	Page      int      `form:"page" binding:"required" example:"1"`
	Offset    int      `swaggerignore:"true"`
	UserTypes []string `form:"user_types"`
	Email     string   `form:"email"`
}

type AdminCreateStudentAccountInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
	SubMajorID  int64  `json:"sub_major_id" binding:"required"`
}

type AdminCreateTeacherAccountInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	SubMajorID  int64  `json:"sub_major_id" binding:"required"`
}

type ListUsersOutput struct {
	Meta  dto.MetaPagination       `json:"meta"`
	Items []user_dto.GetUserOutput `json:"items"`
}

type AccountWithEmail interface {
	GetEmail() string
}

func (input AdminCreateStudentAccountInput) GetEmail() string {
	return input.Email
}

func (input AdminCreateTeacherAccountInput) GetEmail() string {
	return input.Email
}
