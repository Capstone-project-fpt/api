package evaluation_committee_dto

import (
	"time"

	"github.com/api/database/model"
)

type CreateEvaluationCommitteeInput struct {
	Name       string  `json:"name" binding:"required" validate:"required"`
	TeacherIDs []int64 `json:"teacher_ids" binding:"required" validate:"required,min=1"`
	SemesterID int64   `json:"semester_id" binding:"required" validate:"required"`
}

type EvaluationCommitteeOutput struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	TeacherIDs []int64   `json:"teacher_ids"`
	SemesterID int64     `json:"semester_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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

// This used for swagger
type EvaluationCommitteeSwaggerOutput struct {
	Code    int                       `json:"code"`
	Success bool                      `json:"message"`
	Data    EvaluationCommitteeOutput `json:"data"`
}
