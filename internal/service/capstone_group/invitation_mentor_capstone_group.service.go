package capstone_group_service

import (
	"errors"
	"time"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/internal/queue"
	context_util "github.com/api/pkg/utils/context"
	jwt_util "github.com/api/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (cgs *capstoneGroupService) InviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.InviteMentorToCapstoneGroupInput) error {
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
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if capstoneGroup.LeaderID != currentStudent.ID {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	if capstoneGroup.MentorID != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupAlreadyHaveMentor,
		}))
	}

	var teacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Preload("User").Where("id = ?", input.TeacherID).First(&teacher).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var totalCapstoneGroupTeacherMentor int64
	if err := global.Db.Model(model.CapstoneGroup{}).
		Where("mentor_id = ? AND semester_id = ?", input.TeacherID, input.SemesterID).
		Count(&totalCapstoneGroupTeacherMentor).Error; err != nil {
		return err
	}

	if totalCapstoneGroupTeacherMentor >= constant.MaxTotalCapstoneGroupTeacherMentor {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.MaxTotalCapstoneGroupTeacherMentor,
		}))
	}

	tokenExpirationAt := time.Now().Add(time.Duration(constant.DefaultInviteMentorTokenExpiration) * time.Second)

	invitationMentor := model.InvitationMentorCapstoneGroup{
		Status:          constant.InvitationMentorCapstoneGroup.Pending,
		CapstoneGroupID: input.CapstoneGroupID,
		MentorID:        input.TeacherID,
		ExpiredAt:       tokenExpirationAt,
	}
	if err := global.Db.Model(model.InvitationMentorCapstoneGroup{}).Create(&invitationMentor).Error; err != nil {
		return err
	}

	token, err := jwt_util.GenerateInviteMentorToken(
		jwt_util.InviteMentorJwtInput{
			TeacherID:       input.TeacherID,
			CapstoneGroupID: input.CapstoneGroupID,
			InviteID:        invitationMentor.ID,
		},
		tokenExpirationAt.UnixMilli(),
	)
	if err != nil {
		return err
	}

	// TODO: Create Rate Limit to prevent spam
	if err := cgs.emailInviteMentorToCapstoneGroupPublisher.SendMessage(
		queue.InviteMentorToCapstoneGroupMessage{
			MentorID:          int(input.TeacherID),
			MentorEmail:       teacher.User.Email,
			CapstoneGroupID:   int(input.CapstoneGroupID),
			CapstoneGroupName: capstoneGroup.NameGroup,
			Token:             token,
		},
		global.Config.AsynqSetting.DelayInSeconds,
	); err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) ResponseInviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.ResponseInviteMentorToCapstoneGroupInput) error {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	var currentTeacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Where("user_id = ?", currentUser.ID).First(&currentTeacher).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	parseToken, err := jwt_util.VerifyInviteMentorToken(input.Token)
	if err != nil {
		return err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if capstoneGroup.MentorID != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupAlreadyMentor,
		}))
	}

	if capstoneGroup.ID != parseToken.CapstoneGroupID || currentTeacher.ID != parseToken.TeacherID {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	var invitationMentor model.InvitationMentorCapstoneGroup
	if err := global.Db.Model(model.InvitationMentorCapstoneGroup{}).Where("id = ?", parseToken.InviteID).First(&invitationMentor).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvitationMentorCapstoneGroupNotFound,
		}))
	}

	if invitationMentor.ExpiredAt.Before(time.Now()) {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvitationMentorCapstoneGroupExpired,
		}))
	}

	if invitationMentor.Status != constant.InvitationMentorCapstoneGroup.Pending {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvitationMentorCapstoneGroupAlreadyResponded,
			TemplateData: map[string]interface{}{
				"Status": invitationMentor.Status,
			},
		}))
	}

	var totalCapstoneGroupTeacherMentor int64
	if err := global.Db.Model(model.CapstoneGroup{}).
		Where("mentor_id = ? AND semester_id = ?", currentTeacher.ID, capstoneGroup.SemesterID).
		Count(&totalCapstoneGroupTeacherMentor).Error; err != nil {
		return err
	}

	if totalCapstoneGroupTeacherMentor >= constant.MaxTotalCapstoneGroupTeacherMentor && input.Status == constant.InvitationMentorCapstoneGroup.Approve {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.MaxTotalCapstoneGroupTeacherMentor,
		}))
	}

	if err := global.Db.Model(model.InvitationMentorCapstoneGroup{}).Where("id = ?", parseToken.InviteID).Updates(&model.InvitationMentorCapstoneGroup{
		Status: input.Status,
	}).Error; err != nil {
		return err
	}

	if input.Status == constant.InvitationMentorCapstoneGroup.Approve {
		if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).Updates(&model.CapstoneGroup{
			MentorID: &currentTeacher.ID,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

func (cgs *capstoneGroupService) GetListInvitationMentorCapstoneGroups(ctx *gin.Context, input *capstone_group_dto.GetListInviteMentorToCapstoneGroupInput) (*capstone_group_dto.ListInvitationMentorCapstoneGroupOutput, error) {
	var total int64
	var items []model.InvitationMentorCapstoneGroup

	if err := global.Db.Model(model.InvitationMentorCapstoneGroup{}).Where("capstone_group_id = ?", input.CapstoneGroupID).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := global.Db.Model(model.InvitationMentorCapstoneGroup{}).
		Joins("Mentor.User").
		Where("capstone_group_id = ?", input.CapstoneGroupID).
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsInvitationMentorCapstoneGroupOutput := make([]*capstone_group_dto.InvitationMentorCapstoneGroupOutput, len(items))
	for i, item := range items {
		itemsInvitationMentorCapstoneGroupOutput[i] = capstone_group_dto.ToInvitationMentorCapstoneGroupOutput(&item)
	}

	return &capstone_group_dto.ListInvitationMentorCapstoneGroupOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsInvitationMentorCapstoneGroupOutput,
	}, nil
}
