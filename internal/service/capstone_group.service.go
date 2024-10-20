package service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_dto"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

type ICapstoneGroupService interface {
	CreateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupInput) error
	UpdateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupInput) error
	GetCapstoneGroup(ctx *gin.Context, id int) (*capstone_group_dto.CapstoneGroupOutput, error)
	GetListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetListCapstoneGroupInput) (*capstone_group_dto.ListCapstoneGroupOutput, error)
}

type capstoneGroupService struct{}

func NewCapstoneGroupService() ICapstoneGroupService {
	return &capstoneGroupService{}
}

func (cgs *capstoneGroupService) CreateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupInput) error {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var currentStudent model.Student
	if err := global.Db.Model(model.Student{}).Where("user_id = ?", currentUser.ID).First(&currentStudent).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	input.StudentIds = funk.FilterInt64(input.StudentIds, func(id int64) bool {
		return id != currentStudent.ID
	})

	totalMembers := len(input.StudentIds) + 1

	if totalMembers > constant.MaxTotalMemberInGroup || totalMembers < constant.MinTotalMemberInGroup {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidTotalMemberInGroup,
		}))
	}

	var memberGroups []model.Student
	if err := global.Db.Model(model.Student{}).Where("id IN ?", input.StudentIds).Find(&memberGroups).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	if len(memberGroups) != len(input.StudentIds) {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var major model.Major
	if err := global.Db.Model(model.Major{}).Where("id = ?", input.MajorID).First(&major).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.MajorNotFound,
		}))
	}

	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", input.SemesterID).First(&semester).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		}))
	}

	capstoneGroup := model.CapstoneGroup{
		NameGroup:  input.NameGroup,
		MajorID:    input.MajorID,
		SemesterID: input.SemesterID,
		LeaderID:   currentStudent.ID,
	}

	if err := global.Db.Model(model.CapstoneGroup{}).Create(&capstoneGroup).Error; err != nil {
		return err
	}

	if err := global.Db.Model(model.Student{}).
		Where("id IN ?", append(input.StudentIds, currentStudent.ID)).
		Update("capstone_group_id", capstoneGroup.ID).
		Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) UpdateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupInput) error {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	var currentStudent model.Student
	if err := global.Db.Model(model.Student{}).Where("user_id = ?", currentUser.ID).First(&currentStudent).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.ID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if capstoneGroup.LeaderID != currentStudent.ID {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.ID).Updates(&model.CapstoneGroup{
		NameGroup:  input.NameGroup,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) GetCapstoneGroup(ctx *gin.Context, id int) (*capstone_group_dto.CapstoneGroupOutput, error) {
	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", id).First(&capstoneGroup).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupOutput(&capstoneGroup)
	return &capstoneGroupOutput, nil
}

func (cgs *capstoneGroupService) GetListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetListCapstoneGroupInput) (*capstone_group_dto.ListCapstoneGroupOutput, error) {
	var total int64
	var items []model.CapstoneGroup

	if err := global.Db.Model(model.CapstoneGroup{}).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := global.Db.Model(model.CapstoneGroup{}).
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsCapstoneGroupOutput := make([]capstone_group_dto.CapstoneGroupOutput, len(items))
	for i, item := range items {
		itemsCapstoneGroupOutput[i] = capstone_group_dto.ToCapstoneGroupOutput(&item)
	}

	return &capstone_group_dto.ListCapstoneGroupOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsCapstoneGroupOutput,
	}, nil
}
