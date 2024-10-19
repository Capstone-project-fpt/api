package semester_dto

import (
	"time"

	"github.com/api/database/model"
	"github.com/api/internal/dto"
)

type GetListSemestersInput struct {
	Limit  int `form:"limit" binding:"required" example:"10"`
	Page   int `form:"page" binding:"required" example:"1"`
	Offset int `swaggerignore:"true"`
}

type CreateSemesterInput struct {
	Name      string    `json:"name" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required" example:"2006-01-02T15:04:05Z"`
	EndTime   time.Time `json:"end_time" binding:"required" example:"2006-01-02T15:04:05Z"`
}

type UpdateSemesterInput struct {
	ID        int64     `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required" example:"2006-01-02T15:04:05Z"`
	EndTime   time.Time `json:"end_time" binding:"required" example:"2006-01-02T15:04:05Z"`
}

type SemesterOutput struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToSemesterOutput(semester *model.Semester) SemesterOutput {
	return SemesterOutput{
		ID:        semester.ID,
		Name:      semester.Name,
		StartTime: semester.StartTime,
		EndTime:   semester.EndTime,
		CreatedAt: semester.CreatedAt,
		UpdatedAt: semester.UpdatedAt,
	}
}

type ListSemestersOutput struct {
	Meta  dto.MetaPagination `json:"meta"`
	Items []SemesterOutput   `json:"items"`
}

// This used for swagger
type GetSemesterSwaggerOutput struct {
	Code    int             `json:"code"`
	Success bool            `json:"message"`
	Data    *SemesterOutput `json:"data"`
}
