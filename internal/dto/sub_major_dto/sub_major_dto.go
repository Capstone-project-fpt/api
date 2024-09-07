package sub_major_dto

import (
	"github.com/api/database/model"
	"github.com/api/internal/dto"
)

type InputGetListSubMajor struct {
	Limit   int `form:"limit" binding:"required" example:"10"`
	Page    int `form:"page" binding:"required" example:"1"`
	Offset  int `swaggerignore:"true"`
	MajorID int `form:"major_id"`
}

type InputGetSubMajor struct {
	ID int `form:"id" binding:"required" example:"1"`
}

type OutputSubMajor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	MajorID int    `json:"major_id"`
}

type OutputGetListMajor struct {
	Meta  dto.MetaPagination `json:"meta"`
	Items []OutputSubMajor   `json:"items"`
}

func ToSubMajorOutput(subMajor model.SubMajor) OutputSubMajor {
	return OutputSubMajor{
		ID:      int(subMajor.ID),
		Name:    subMajor.Name,
		MajorID: int(subMajor.MajorID),
	}
}

// This is used for swagger
type OutputGetSubMajorSwagger struct {
	Code    int            `json:"code"`
	Success bool           `json:"message"`
	Data    OutputSubMajor `json:"data"`
}
