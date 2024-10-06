package major_dto

import (
	"github.com/api/database/model"
	"github.com/api/internal/dto"
)

type GetListMajorInput struct {
	Limit  int `form:"limit" binding:"required" example:"10"`
	Page   int `form:"page" binding:"required" example:"1"`
	Offset int `swaggerignore:"true"`
}

type MajorOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
}

type GetListMajorOutput struct {
	Meta  dto.MetaPagination `json:"meta"`
	Items []MajorOutput      `json:"items"`
}

func ToMajorOutput(major model.Major) MajorOutput {
	return MajorOutput{
		ID:        int(major.ID),
		Name:      major.Name,
	}
}

// This is used for swagger
type GetMajorSwaggerOutput struct {
	Code    int         `json:"code"`
	Success bool        `json:"message"`
	Data    MajorOutput `json:"data"`
}