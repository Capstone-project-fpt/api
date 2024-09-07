package major_dto

import (
	"github.com/api/database/model"
	"github.com/api/internal/dto"
)

type InputGetListMajor struct {
	Limit  int `form:"limit" binding:"required" example:"10"`
	Page   int `form:"page" binding:"required" example:"1"`
	Offset int `swaggerignore:"true"`
}

type OutputMajor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
}

type OutputGetListMajor struct {
	Meta  dto.MetaPagination `json:"meta"`
	Items []OutputMajor      `json:"items"`
}

func ToMajorOutput(major model.Major) OutputMajor {
	return OutputMajor{
		ID:        int(major.ID),
		Name:      major.Name,
	}
}
