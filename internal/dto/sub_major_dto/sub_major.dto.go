package sub_major_dto

import (
	"github.com/api/database/model"
	"github.com/api/internal/dto"
)

type GetListSubMajorInput struct {
	Limit   int `form:"limit" binding:"required" example:"10"`
	Page    int `form:"page" binding:"required" example:"1"`
	Offset  int `swaggerignore:"true"`
	MajorID int `form:"major_id"`
}

type GetSubMajorInput struct {
	ID int `form:"id" binding:"required" example:"1"`
}

type SubMajorOutput struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	MajorID int    `json:"major_id"`
}

type GetListSubMajorOutput struct {
	Meta  dto.MetaPagination `json:"meta"`
	Items []SubMajorOutput   `json:"items"`
}

func ToSubMajorOutput(subMajor model.SubMajor) SubMajorOutput {
	return SubMajorOutput{
		ID:      int(subMajor.ID),
		Name:    subMajor.Name,
		MajorID: int(subMajor.MajorID),
	}
}

// This is used for swagger
type GetSubMajorSwaggerOutput struct {
	Code    int            `json:"code"`
	Success bool           `json:"message"`
	Data    SubMajorOutput `json:"data"`
}
