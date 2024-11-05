package evaluation_committee_dto

import (
	"time"

	"github.com/api/database/model"
)

type EvaluationCommitteeOutput struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	TeacherIds []int64   `json:"teacher_ids"`
	SemesterID int64     `json:"semester_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToEvaluationCommitteeOutput(e *model.EvaluationCommittee) *EvaluationCommitteeOutput {
	return &EvaluationCommitteeOutput{
		ID:         e.ID,
		Name:       e.Name,
		TeacherIds: e.TeacherIds,
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
