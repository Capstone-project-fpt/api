package topic_reference_dto

import (
	"github.com/api/database/model"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/user_dto"
)

type GetListTopicReferencesInput struct {
	Limit      int    `form:"limit" binding:"required" example:"10"`
	Page       int    `form:"page" binding:"required" example:"1"`
	Offset     int    `swaggerignore:"true"`
	TeacherIDs []int  `form:"teacher_ids"`
	Search     string `form:"search"`
}

type TeacherCreateTopicReferenceInput struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

type TeacherUpdateTopicReferenceInput struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

type TeacherDeleteTopicReferenceInput struct {
	ID int `form:"id" binding:"required" example:"1"`
}

type AdminCreateTopicReferenceInput struct {
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path" binding:"required"`
	TeacherID int64  `json:"teacher_id" binding:"required"`
}

type TopicReferenceOutput struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	Path    string                 `json:"path"`
	Teacher user_dto.TeacherOutput `json:"teacher"`
}

func ToTopicReferenceOutput(topicReference *model.TopicReferences) TopicReferenceOutput {
	return TopicReferenceOutput{
		ID:      int(topicReference.ID),
		Name:    topicReference.Name,
		Path:    topicReference.Path,
		Teacher: user_dto.ToTeacherOutput(&topicReference.Teacher),
	}
}

type ListTopicReferenceOutput struct {
	Meta  dto.MetaPagination     `json:"meta"`
	Items []TopicReferenceOutput `json:"items"`
}

// This used for swagger
type GetTopicReferenceSwaggerOutput struct {
	Code    int                   `json:"code"`
	Success bool                  `json:"message"`
	Data    *TopicReferenceOutput `json:"data"`
}
