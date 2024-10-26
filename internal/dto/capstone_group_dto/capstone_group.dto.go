package capstone_group_dto

import (
	"time"

	"github.com/api/database/model"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/user_dto"
)

type CreateCapstoneGroupInput struct {
	NameGroup  string  `json:"name_group" binding:"required"`
	StudentIds []int64 `json:"student_ids" binding:"required"`
	SemesterID int64   `json:"semester_id" binding:"required"`
	MajorID    int64   `json:"major_id" binding:"required"`
}

type UpdateCapstoneGroupInput struct {
	ID        int64  `json:"id" binding:"required"`
	NameGroup string `json:"name_group" binding:"required"`
}

type InviteMentorToCapstoneGroupInput struct {
	TeacherID       int64 `json:"teacher_id" binding:"required"`
	SemesterID      int64 `json:"semester_id" binding:"required"`
	CapstoneGroupID int64 `swaggerignore:"true"`
}

type ResponseInviteMentorToCapstoneGroupInput struct {
	Token           string `json:"token" binding:"required"`
	Status          string `json:"status" binding:"required" validate:"required,oneof=approve reject"`
	CapstoneGroupID int64  `swaggerignore:"true"`
}

type GetListCapstoneGroupInput struct {
	Limit  int `form:"limit" binding:"required" example:"10"`
	Page   int `form:"page" binding:"required" example:"1"`
	Offset int `swaggerignore:"true"`
}

type GetListInviteMentorToCapstoneGroupInput struct {
	Limit           int   `form:"limit" binding:"required" example:"10"`
	Page            int   `form:"page" binding:"required" example:"1"`
	Offset          int   `swaggerignore:"true"`
	CapstoneGroupID int64 `swaggerignore:"true"`
}

type CapstoneGroupOutput struct {
	ID         int64     `json:"id"`
	NameGroup  string    `json:"name_group"`
	TopicID    *int64    `json:"topic_id"`
	MajorID    int64     `json:"major_id"`
	SemesterID int64     `json:"semester_id"`
	LeaderID   int64     `json:"leader_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToCapstoneGroupOutput(capstoneGroup *model.CapstoneGroup) CapstoneGroupOutput {
	return CapstoneGroupOutput{
		ID:         capstoneGroup.ID,
		NameGroup:  capstoneGroup.NameGroup,
		TopicID:    capstoneGroup.TopicID,
		MajorID:    capstoneGroup.MajorID,
		SemesterID: capstoneGroup.SemesterID,
		LeaderID:   capstoneGroup.LeaderID,
		Status:     capstoneGroup.Status,
		CreatedAt:  capstoneGroup.CreatedAt,
		UpdatedAt:  capstoneGroup.UpdatedAt,
	}
}

type InvitationMentorCapstoneGroupOutput struct {
	ID              int64                   `json:"id"`
	Status          string                  `json:"status"`
	CapstoneGroupID int64                   `json:"capstone_group_id"`
	MentorID        int64                   `json:"mentor_id"`
	Mentor          *user_dto.TeacherOutput `json:"mentor"`
	ExpiredAt       time.Time               `json:"expired_at"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}

func ToInvitationMentorCapstoneGroupOutput(invitationMentorCapstoneGroup *model.InvitationMentorCapstoneGroup) *InvitationMentorCapstoneGroupOutput {
	return &InvitationMentorCapstoneGroupOutput{
		ID:              invitationMentorCapstoneGroup.ID,
		Status:          invitationMentorCapstoneGroup.Status,
		CapstoneGroupID: invitationMentorCapstoneGroup.CapstoneGroupID,
		MentorID:        invitationMentorCapstoneGroup.MentorID,
		Mentor:          user_dto.ToTeacherOutput(&invitationMentorCapstoneGroup.Mentor),
		ExpiredAt:       invitationMentorCapstoneGroup.ExpiredAt,
		CreatedAt:       invitationMentorCapstoneGroup.CreatedAt,
		UpdatedAt:       invitationMentorCapstoneGroup.UpdatedAt,
	}
}

type ListInvitationMentorCapstoneGroupOutput struct {
	Meta  dto.MetaPagination                     `json:"meta"`
	Items []*InvitationMentorCapstoneGroupOutput `json:"items"`
}

type ListCapstoneGroupOutput struct {
	Meta  dto.MetaPagination    `json:"meta"`
	Items []CapstoneGroupOutput `json:"items"`
}

type MentorAndListMemberCapstoneGroupOutput struct {
	Mentor   *user_dto.TeacherOutput   `json:"mentor"`
	Members  []*user_dto.StudentOutput `json:"members"`
	LeaderID int64                     `json:"leader_id"`
}

// This used for swagger
type GetCapstoneGroupSwaggerOutput struct {
	Code    int                  `json:"code"`
	Success bool                 `json:"message"`
	Data    *CapstoneGroupOutput `json:"data"`
}

type MentorAndListMemberCapstoneGroupSwaggerOutput struct {
	Code    int                                     `json:"code"`
	Success bool                                    `json:"message"`
	Data    *MentorAndListMemberCapstoneGroupOutput `json:"data"`
}
