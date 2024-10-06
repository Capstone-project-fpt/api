package service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/topic_reference_dto"
	"github.com/api/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ITopicReferenceService interface {
	GetTopicReference(ctx *gin.Context, id int) (*topic_reference_dto.TopicReferenceOutput, error)
	GetListTopicReferences(ctx *gin.Context, input *topic_reference_dto.GetListTopicReferencesInput) (interface{}, error)
	TeacherCreateTopicReference(ctx *gin.Context, input CreateTopicReferenceInput, userContext *types.UserContext) error
	AdminCreateTopicReference(ctx *gin.Context, input CreateTopicReferenceInput) error
	UpdateTopicReference(ctx *gin.Context, input UpdateTopicReferenceInput) error
	DeleteTopicReference(ctx *gin.Context, input DeleteTopicReferenceInput) error
}

type topicReferenceService struct{}

func NewTopicReferenceService() ITopicReferenceService {
	return &topicReferenceService{}
}

type CreateTopicReferenceInput struct {
	Name      string
	TeacherID int64
	Path      string
}

type UpdateTopicReferenceInput struct {
	ID        int64
	Name      string
	TeacherID int64
	Path      string
}

type DeleteTopicReferenceInput struct {
	ID        int64
	TeacherID int64
}

func (tr *topicReferenceService) GetTopicReference(ctx *gin.Context, id int) (*topic_reference_dto.TopicReferenceOutput, error) {
	var topicReference model.TopicReferences
	err := global.Db.Joins("Teacher.User").Where("topic_references.id = ?", id).First(&topicReference).Error
	if err != nil {
		return nil, err
	}

	output := topic_reference_dto.ToTopicReferenceOutput(&topicReference)

	return &output, nil
}

func (tr *topicReferenceService) GetListTopicReferences(ctx *gin.Context, input *topic_reference_dto.GetListTopicReferencesInput) (interface{}, error) {
	var total int64
	var items []model.TopicReferences
	getTotalQuery := global.Db.Model(&model.TopicReferences{})
	getTopicReferencesQuery := global.Db.Model(&model.TopicReferences{}).Joins("Teacher.User")

	if input.TeacherID != 0 {
		getTotalQuery = getTotalQuery.Where("teacher_id = ?", input.TeacherID)
		getTopicReferencesQuery = getTopicReferencesQuery.Where("topic_references.teacher_id = ?", input.TeacherID)
	}

	if err := getTotalQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := getTopicReferencesQuery.Omit("CreatedAt", "UpdatedAt").
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsTopicReferenceOutput := make([]topic_reference_dto.TopicReferenceOutput, len(items))
	for i, item := range items {
		itemsTopicReferenceOutput[i] = topic_reference_dto.ToTopicReferenceOutput(&item)
	}

	return topic_reference_dto.ListTopicReferenceOutput{
		Meta:  dto.MetaPagination{
			CurrentPage: int(input.Page), 
			Total: int(total),
		},
		Items: itemsTopicReferenceOutput,
	}, nil
}

func (tr *topicReferenceService) UpdateTopicReference(ctx *gin.Context, input UpdateTopicReferenceInput) error {
	var topicReference model.TopicReferences
	err := global.Db.Model(model.TopicReferences{}).
		Where("id = ? AND teacher_id = ?", input.ID, input.TeacherID).
		First(&topicReference).
		Error
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TopicReferenceNotFound,
		})
		return errors.New(message)
	}

	topicReference.Name = input.Name
	topicReference.Path = input.Path
	global.Db.Save(&topicReference)

	return nil
}

func (tr *topicReferenceService) TeacherCreateTopicReference(ctx *gin.Context, input CreateTopicReferenceInput, userContext *types.UserContext) error {
	var teacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Select("id").Where("user_id = ?", userContext.ID).First(&teacher).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return errors.New(message)
	}

	input.TeacherID = teacher.ID

	return tr.CreateTopicReference(ctx, input)
}

func (tr *topicReferenceService) DeleteTopicReference(ctx *gin.Context, input DeleteTopicReferenceInput) error {
	var topicReference model.TopicReferences
	err := global.Db.Model(model.TopicReferences{}).
		Where("id = ?", input.ID).
		First(&topicReference).
		Error
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TopicReferenceNotFound,
		})
		return errors.New(message)
	}
	if topicReference.TeacherID != input.TeacherID {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		})
		return errors.New(message)
	}

	global.Db.Delete(&topicReference)

	return nil
}

func (tr *topicReferenceService) AdminCreateTopicReference(ctx *gin.Context, input CreateTopicReferenceInput) error {

	var teacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Select("id").Where("id = ?", input.TeacherID).First(&teacher).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return errors.New(message)
	}

	input.TeacherID = teacher.ID

	return tr.CreateTopicReference(ctx, input)
}

func (tr *topicReferenceService) CreateTopicReference(ctx *gin.Context, input CreateTopicReferenceInput) error {
	topicReference := model.TopicReferences{
		Name:      input.Name,
		Path:      input.Path,
		TeacherID: input.TeacherID,
	}

	err := global.Db.Create(&topicReference).Error
	if err != nil {
		return err
	}

	return nil
}
