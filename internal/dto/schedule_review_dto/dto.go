package schedule_review_dto

import (
	"time"

	"github.com/api/database/model"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/internal/dto/evaluation_committee_dto"
)

type CreateScheduleReviewInput struct {
	Title                 string    `json:"title" validate:"required"`
	Description           string    `json:"description" validate:"required"`
	Type                  string    `json:"type" validate:"required,oneof=first_review second_review third_review"`
	StartTime             time.Time `json:"start_time" validate:"required"`
	EndTime               time.Time `json:"end_time" validate:"required"`
	SemesterID            int64     `json:"semester_id" validate:"required"`
	EvaluationCommitteeID int64     `json:"evaluation_committee_id" validate:"required"`
	CapstoneGroupID       int64     `json:"capstone_group_id" validate:"required"`
}

type UpdateScheduleReviewInput struct {
	ID                    int64     `json:"id" validate:"required"`
	Title                 string    `json:"title" validate:"required"`
	Description           string    `json:"description" validate:"required"`
	Type                  string    `json:"type" validate:"required,oneof=first_review second_review third_review"`
	StartTime             time.Time `json:"start_time" validate:"required"`
	EndTime               time.Time `json:"end_time" validate:"required"`
	SemesterID            int64     `json:"semester_id" validate:"required"`
	EvaluationCommitteeID int64     `json:"evaluation_committee_id" validate:"required"`
	CapstoneGroupID       int64     `json:"capstone_group_id" validate:"required"`
}

type GetListScheduleReviewInput struct {
	StartTime             time.Time `form:"start_time" binding:"required" example:"2024-06-01T08:00:00Z"`
	EndTime               time.Time `form:"end_time" binding:"required" example:"2024-06-01T08:00:00Z"`
	OrderBy               *string   `form:"order_by" example:"ASC|DESC"`
	EvaluationCommitteeID *int64    `form:"evaluation_committee_id" example:"1"`
	CapstoneGroupID       *int64    `form:"capstone_group_id" example:"1"`
}

type ScheduleReviewOutput struct {
	ID                    int64     `json:"id"`
	Title                 string    `json:"title"`
	Description           string    `json:"description"`
	LinkMeeting           string    `json:"link_meeting"`
	Type                  string    `json:"type"`
	StartTime             time.Time `json:"start_time"`
	EndTime               time.Time `json:"end_time"`
	EvaluationCommitteeID int64     `json:"evaluation_committee_id"`
	CapstoneGroupID       int64     `json:"capstone_group_id"`
}

type ScheduleReviewDetailOutput struct {
	ID                  int64                                                              `json:"id"`
	Title               string                                                             `json:"title"`
	Description         string                                                             `json:"description"`
	LinkMeeting         string                                                             `json:"link_meeting"`
	Type                string                                                             `json:"type"`
	StartTime           time.Time                                                          `json:"start_time"`
	EndTime             time.Time                                                          `json:"end_time"`
	EvaluationCommittee *evaluation_committee_dto.EvaluationCommitteeWithTeacherInfoOutput `json:"evaluation_committee"`
	CapstoneGroup       *capstone_group_dto.CapstoneGroupOutput                            `json:"capstone_group"`
	CapstoneGroupReview *capstone_group_dto.CapstoneGroupReviewOutput                      `json:"capstone_group_review"`
}

func ToScheduleReviewDetailOutput(
	scheduleReview *model.ScheduleReview,
	evaluationCommittee *evaluation_committee_dto.EvaluationCommitteeWithTeacherInfoOutput,
	capstoneGroup *capstone_group_dto.CapstoneGroupOutput,
	capstoneGroupReview *capstone_group_dto.CapstoneGroupReviewOutput,
) *ScheduleReviewDetailOutput {
	return &ScheduleReviewDetailOutput{
		ID:                  scheduleReview.ID,
		Title:               scheduleReview.Title,
		Description:         scheduleReview.Description,
		LinkMeeting:         scheduleReview.LinkMeeting,
		Type:                scheduleReview.Type,
		StartTime:           scheduleReview.StartTime,
		EndTime:             scheduleReview.EndTime,
		EvaluationCommittee: evaluationCommittee,
		CapstoneGroup:       capstoneGroup,
		CapstoneGroupReview: capstoneGroupReview,
	}
}

func ToScheduleReviewOutput(scheduleReview *model.ScheduleReview) *ScheduleReviewOutput {
	return &ScheduleReviewOutput{
		ID:                    scheduleReview.ID,
		Title:                 scheduleReview.Title,
		Description:           scheduleReview.Description,
		LinkMeeting:           scheduleReview.LinkMeeting,
		Type:                  scheduleReview.Type,
		StartTime:             scheduleReview.StartTime,
		EndTime:               scheduleReview.EndTime,
		EvaluationCommitteeID: scheduleReview.EvaluationCommitteeID,
		CapstoneGroupID:       scheduleReview.CapstoneGroupID,
	}
}

// This used for swagger
type ScheduleReviewSwaggerOutput struct {
	Code    int                  `json:"code"`
	Success bool                 `json:"message"`
	Data    ScheduleReviewOutput `json:"data"`
}

type ScheduleReviewDetailSwaggerOutput struct {
	Code    int                        `json:"code"`
	Success bool                       `json:"message"`
	Data    ScheduleReviewDetailOutput `json:"data"`
}

type ListScheduleReviewSwaggerOutput struct {
	Code    int                    `json:"code"`
	Success bool                   `json:"message"`
	Data    []ScheduleReviewOutput `json:"data"`
}

type ListScheduleReviewDetailSwaggerOutput struct {
	Code    int                          `json:"code"`
	Success bool                         `json:"message"`
	Data    []ScheduleReviewDetailOutput `json:"data"`
}
