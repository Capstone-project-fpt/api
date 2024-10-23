package capstone_group_topic_dto

import (
	"time"

	"github.com/api/database/model"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/user_dto"
)

type CreateCapstoneGroupTopicInput struct {
	Topic           string `json:"topic"`
	DocumentPath    string `json:"document_path"`
	CapstoneGroupID int64  `swaggerignore:"true"`
}

type UpdateCapstoneGroupTopicInput struct {
	Topic                string `json:"topic"`
	DocumentPath         string `json:"document_path"`
	CapstoneGroupID      int64  `swaggerignore:"true"`
	CapstoneGroupTopicID int64  `swaggerignore:"true"`
}

type GetListCapstoneGroupTopicInput struct {
	Limit           int    `form:"limit" binding:"required" example:"10"`
	Page            int    `form:"page" binding:"required" example:"1"`
	OrderBy         string `form:"order_by" example:"ASC|DESC"`
	Offset          int    `swaggerignore:"true"`
	CapstoneGroupID int64  `swaggerignore:"true"`
}

type CapstoneGroupTopicOutput struct {
	ID              int64                   `json:"id"`
	Topic           string                  `json:"topic"`
	DocumentPath    string                  `json:"document_path"`
	StatusReview    string                  `json:"status_review"`
	ApprovedAt      *time.Time              `json:"approved_at"`
	ApprovedByID    *int64                  `json:"approved_by_id"`
	ApprovedBy      *user_dto.TeacherOutput `json:"approved_by"`
	RejectedAt      *time.Time              `json:"rejected_at"`
	RejectedByID    *int64                  `json:"rejected_by_id"`
	RejectedBy      *user_dto.TeacherOutput `json:"rejected_by"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	CapstoneGroupID int64                   `json:"capstone_group_id"`
}

func ToCapstoneGroupTopicOutput(capstoneGroupTopic *model.CapstoneGroupTopic) CapstoneGroupTopicOutput {
	var approvedBy user_dto.TeacherOutput
	var rejectedBy user_dto.TeacherOutput

	if capstoneGroupTopic.ApprovedBy != nil {
		approvedBy = user_dto.ToTeacherOutput(capstoneGroupTopic.ApprovedBy)
	}

	if capstoneGroupTopic.RejectedBy != nil {
		rejectedBy = user_dto.ToTeacherOutput(capstoneGroupTopic.RejectedBy)
	}

	return CapstoneGroupTopicOutput{
		ID:              capstoneGroupTopic.ID,
		Topic:           capstoneGroupTopic.Topic,
		DocumentPath:    capstoneGroupTopic.DocumentPath,
		StatusReview:    capstoneGroupTopic.StatusReview,
		ApprovedAt:      capstoneGroupTopic.ApprovedAt,
		ApprovedByID:    capstoneGroupTopic.ApprovedByID,
		ApprovedBy:      &approvedBy,
		RejectedAt:      capstoneGroupTopic.RejectedAt,
		RejectedByID:    capstoneGroupTopic.RejectedByID,
		RejectedBy:      &rejectedBy,
		CreatedAt:       capstoneGroupTopic.CreatedAt,
		UpdatedAt:       capstoneGroupTopic.UpdatedAt,
		CapstoneGroupID: capstoneGroupTopic.CapstoneGroupID,
	}
}

type ListCapstoneGroupTopicsOutput struct {
	Meta  dto.MetaPagination         `json:"meta"`
	Items []CapstoneGroupTopicOutput `json:"items"`
}

// This used for swagger
type GetCapstoneTopicGroupSwaggerOutput struct {
	Code    int                       `json:"code"`
	Success bool                      `json:"message"`
	Data    *CapstoneGroupTopicOutput `json:"data"`
}
