package evaluation_committee_dto

import (
	"time"

	"github.com/api/database/model"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/user_dto"
)

type CreateEvaluationCommitteeInput struct {
	Name       string  `json:"name" binding:"required" validate:"required"`
	TeacherIDs []int64 `json:"teacher_ids" binding:"required" validate:"required,min=1"`
	SemesterID int64   `json:"semester_id" binding:"required" validate:"required"`
}

type UpdateEvaluationCommitteeInput struct {
	ID         int64   `json:"id" binding:"required" validate:"required"`
	Name       string  `json:"name" binding:"required" validate:"required"`
	TeacherIDs []int64 `json:"teacher_ids" binding:"required" validate:"required,min=1"`
}

type GetListEvaluationCommitteeInput struct {
	Limit      int     `form:"limit" binding:"required" example:"10"`
	Page       int     `form:"page" binding:"required" example:"1"`
	OrderBy    *string `form:"order_by" example:"ASC|DESC"`
	SemesterID *int64  `form:"semester_id" example:"1"`
	Offset     int     `swaggerignore:"true"`
}

type EvaluationCommitteeOutput struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	TeacherIDs []int64   `json:"teacher_ids"`
	SemesterID int64     `json:"semester_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EvaluationCommitteeWithTeacherInfoOutput struct {
	ID         int64                      `json:"id"`
	Name       string                     `json:"name"`
	Teachers   *[]*user_dto.TeacherOutput `json:"teachers"`
	SemesterID int64                      `json:"semester_id"`
	CreatedAt  time.Time                  `json:"created_at"`
	UpdatedAt  time.Time                  `json:"updated_at"`
}

func ToEvaluationCommitteeOutput(e *model.EvaluationCommittee) *EvaluationCommitteeOutput {
	return &EvaluationCommitteeOutput{
		ID:         e.ID,
		Name:       e.Name,
		TeacherIDs: e.TeacherIDs,
		SemesterID: e.SemesterID,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}

func ToEvaluationCommitteeWithTeacherInfoOutput(e *model.EvaluationCommittee, teachers *[]*user_dto.TeacherOutput) *EvaluationCommitteeWithTeacherInfoOutput {
	return &EvaluationCommitteeWithTeacherInfoOutput{
		ID:         e.ID,
		Name:       e.Name,
		Teachers:   teachers,
		SemesterID: e.SemesterID,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}

type ListEvaluationCommitteeOutput struct {
	Meta  dto.MetaPagination           `json:"meta"`
	Items []*EvaluationCommitteeOutput `json:"items"`
}

// This used for swagger
type EvaluationCommitteeSwaggerOutput struct {
	Code    int                       `json:"code"`
	Success bool                      `json:"message"`
	Data    EvaluationCommitteeOutput `json:"data"`
}

type EvaluationCommitteeWithTeacherInfoSwaggerOutput struct {
	Code    int                                      `json:"code"`
	Success bool                                     `json:"message"`
	Data    EvaluationCommitteeWithTeacherInfoOutput `json:"data"`
}
type ListTeachersHaveEvaluationCommitteeGroupSwaggerOutput struct {
	Code    int                        `json:"code"`
	Success bool                       `json:"message"`
	Data    *[]*user_dto.TeacherOutput `json:"data"`
}
