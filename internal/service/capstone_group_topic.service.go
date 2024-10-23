package service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_topic_dto"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ICapstoneGroupTopicService interface {
	CreateCapstoneGroupTopic(ctx *gin.Context, input *capstone_group_topic_dto.CreateCapstoneGroupTopicInput) error
	UpdateCapstoneGroupTopic(ctx *gin.Context, input *capstone_group_topic_dto.UpdateCapstoneGroupTopicInput) error
	DeleteCapstoneGroupTopic(ctx *gin.Context, input *DeleteCapstoneGroupTopicInput) error
	GetCapstoneGroupTopic(ctx *gin.Context, input *GetCapstoneGroupTopicInput) (*capstone_group_topic_dto.CapstoneGroupTopicOutput, error)
	GetListCapstoneGroupTopics(ctx *gin.Context, input *capstone_group_topic_dto.GetListCapstoneGroupTopicInput) (*capstone_group_topic_dto.ListCapstoneGroupTopicsOutput, error)
}

type capstoneGroupTopicService struct{}

func NewCapstoneGroupTopicService() ICapstoneGroupTopicService {
	return &capstoneGroupTopicService{}
}

type DeleteCapstoneGroupTopicInput struct {
	CapstoneGroupTopicID int64
	CapstoneGroupID      int64
}

type GetCapstoneGroupTopicInput struct {
	CapstoneGroupTopicID int64
	CapstoneGroupID      int64
}

func (cgts *capstoneGroupTopicService) CreateCapstoneGroupTopic(ctx *gin.Context, input *capstone_group_topic_dto.CreateCapstoneGroupTopicInput) error {
	currentStudent, err := cgts.getCurrentStudent(ctx)
	if err != nil {
		return err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if err := cgts.validatePermissionActionToCapstoneGroupTopic(currentStudent, &capstoneGroup); err != nil {
		return err
	}

	capstoneGroupTopic := model.CapstoneGroupTopic{
		Topic:           input.Topic,
		DocumentPath:    input.DocumentPath,
		StatusReview:    constant.TopicStatusReview.Reviewing,
		CapstoneGroupID: input.CapstoneGroupID,
	}

	if err := global.Db.Model(model.CapstoneGroupTopic{}).Create(&capstoneGroupTopic).Error; err != nil {
		return err
	}

	return nil
}

func (cgts *capstoneGroupTopicService) UpdateCapstoneGroupTopic(ctx *gin.Context, input *capstone_group_topic_dto.UpdateCapstoneGroupTopicInput) error {
	currentStudent, err := cgts.getCurrentStudent(ctx)
	if err != nil {
		return err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if err := cgts.validatePermissionActionToCapstoneGroupTopic(currentStudent, &capstoneGroup); err != nil {
		return err
	}

	if err := global.Db.Model(&model.CapstoneGroupTopic{}).Where("id = ?", input.CapstoneGroupTopicID).Updates(&model.CapstoneGroupTopic{
		Topic:           input.Topic,
		DocumentPath:    input.DocumentPath,
		StatusReview:    constant.TopicStatusReview.Reviewing,
		CapstoneGroupID: input.CapstoneGroupID,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (cgts *capstoneGroupTopicService) DeleteCapstoneGroupTopic(ctx *gin.Context, input *DeleteCapstoneGroupTopicInput) error {
	currentStudent, err := cgts.getCurrentStudent(ctx)
	if err != nil {
		return err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if err := cgts.validatePermissionActionToCapstoneGroupTopic(currentStudent, &capstoneGroup); err != nil {
		return err
	}

	if err := global.Db.Model(&model.CapstoneGroupTopic{}).Where("id = ?", input.CapstoneGroupTopicID).Delete(&model.CapstoneGroupTopic{}).Error; err != nil {
		return err
	}

	return nil
}

func (cgts *capstoneGroupTopicService) GetCapstoneGroupTopic(ctx *gin.Context, input *GetCapstoneGroupTopicInput) (*capstone_group_topic_dto.CapstoneGroupTopicOutput, error) {
	var capstoneGroupTopic model.CapstoneGroupTopic
	if err := global.Db.Model(model.CapstoneGroupTopic{}).
		Joins("ApprovedBy.User").
		Joins("RejectedBy.User").
		Where("capstone_group_topics.id = ? AND capstone_group_id = ?", input.CapstoneGroupTopicID, input.CapstoneGroupID).
		First(&capstoneGroupTopic).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupTopicNotFound,
		}))
	}

	capstoneGroupTopicOutput := capstone_group_topic_dto.ToCapstoneGroupTopicOutput(&capstoneGroupTopic)

	return &capstoneGroupTopicOutput, nil
}

func (cgts *capstoneGroupTopicService) GetListCapstoneGroupTopics(ctx *gin.Context, input *capstone_group_topic_dto.GetListCapstoneGroupTopicInput) (*capstone_group_topic_dto.ListCapstoneGroupTopicsOutput, error) {
	var total int64
	var items []model.CapstoneGroupTopic

	query := global.Db.Model(model.CapstoneGroupTopic{}).
		Joins("ApprovedBy.User").
		Joins("RejectedBy.User").
		Where("capstone_group_id = ?", input.CapstoneGroupID).
		Limit(int(input.Limit)).
		Offset(int(input.Offset))

	if err := global.Db.Model(model.CapstoneGroupTopic{}).Count(&total).Error; err != nil {
		return nil, err
	}

	if input.OrderBy != "" {
		query = query.Order("capstone_group_topics.created_at " + input.OrderBy)
	}

	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}

	itemsCapstoneGroupTopicOutput := make([]capstone_group_topic_dto.CapstoneGroupTopicOutput, len(items))
	for i, item := range items {
		itemsCapstoneGroupTopicOutput[i] = capstone_group_topic_dto.ToCapstoneGroupTopicOutput(&item)
	}

	return &capstone_group_topic_dto.ListCapstoneGroupTopicsOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsCapstoneGroupTopicOutput,
	}, nil
}

func (cgts *capstoneGroupTopicService) validatePermissionActionToCapstoneGroupTopic(currentStudent *model.Student, capstoneGroup *model.CapstoneGroup) error {
	if currentStudent.CapstoneGroupID != capstoneGroup.ID {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	if capstoneGroup.Status == constant.CapstoneGroupStatus.InProgress {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupInProgress,
		}))
	}

	return nil
}

func (cgts *capstoneGroupTopicService) getCurrentStudent(ctx *gin.Context) (*model.Student, error) {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var currentStudent model.Student
	if err := global.Db.Model(model.Student{}).Where("user_id = ?", currentUser.ID).First(&currentStudent).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	currentStudent.User = model.User{
		ID:       currentUser.ID,
		Email:    currentUser.Email,
		Name:     currentUser.Name,
		UserType: currentUser.UserType,
	}

	return &currentStudent, nil
}
