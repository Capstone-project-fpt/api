package capstone_group_service

import (
	"errors"
	"strings"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/internal/dto/user_dto"
	"github.com/api/internal/queue"
	context_util "github.com/api/pkg/utils/context"
	jwt_util "github.com/api/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

type ICapstoneGroupService interface {
	CreateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupInput) (*capstone_group_dto.CapstoneGroupOutput, error)
	UpdateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupInput) error
	InviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.InviteMentorToCapstoneGroupInput) error
	AcceptInviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.AcceptInviteMentorToCapstoneGroupInput) error
	GetCapstoneGroup(ctx *gin.Context, id int) (*capstone_group_dto.CapstoneGroupOutput, error)
	GetListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetListCapstoneGroupInput) (*capstone_group_dto.ListCapstoneGroupOutput, error)
	GetMentorAndListMemberCapstoneGroup(ctx *gin.Context, id int64) (*capstone_group_dto.MentorAndListMemberCapstoneGroupOutput, error)
}

type capstoneGroupService struct {
	emailInviteMentorToCapstoneGroupPublisher queue.IBasePublisher[queue.InviteMentorToCapstoneGroupMessage]
}

func NewCapstoneGroupService(emailInviteMentorToCapstoneGroupPublisher queue.IBasePublisher[queue.InviteMentorToCapstoneGroupMessage]) ICapstoneGroupService {
	return &capstoneGroupService{
		emailInviteMentorToCapstoneGroupPublisher: emailInviteMentorToCapstoneGroupPublisher,
	}
}

func (cgs *capstoneGroupService) CreateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupInput) (*capstone_group_dto.CapstoneGroupOutput, error) {
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

	input.StudentIds = funk.FilterInt64(input.StudentIds, func(id int64) bool {
		return id != currentStudent.ID
	})

	memberGroupsIDs := append(input.StudentIds, currentStudent.ID)
	totalMembers := len(input.StudentIds) + 1

	if totalMembers > constant.MaxTotalMemberInGroup || totalMembers < constant.MinTotalMemberInGroup {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidTotalMemberInGroup,
		}))
	}

	var memberGroups []model.Student
	if err := global.Db.Model(model.Student{}).Joins("User").Where("students.id IN ?", memberGroupsIDs).Find(&memberGroups).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var memberExistInAnotherGroup []model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Where("student_id IN ? AND semester_id = ?", memberGroupsIDs, input.SemesterID).Find(&memberExistInAnotherGroup).Error; err == nil {
		if len(memberExistInAnotherGroup) > 0 {
			memberIDExistInAnotherGroup := funk.Map(memberExistInAnotherGroup, func(member model.StudentCapstoneGroup) int64 {
				return member.StudentID
			})

			studentExist := funk.Filter(memberGroups, func(student model.Student) bool {
				return funk.ContainsInt64((memberIDExistInAnotherGroup).([]int64), student.ID)
			})

			memberNameArr := funk.Map(studentExist, func(student model.Student) string {
				return student.User.Name
			}).([]string)

			memberNames := strings.Join(memberNameArr, ", ")

			return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.MemberExistInAnotherGroup,
				TemplateData: map[string]interface{}{
					"MemberNames": memberNames,
				},
			}))
		}
	}

	if len(memberGroups) != len(input.StudentIds) {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var major model.Major
	if err := global.Db.Model(model.Major{}).Where("id = ?", input.MajorID).First(&major).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.MajorNotFound,
		}))
	}

	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", input.SemesterID).First(&semester).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		}))
	}

	data := model.CapstoneGroup{
		NameGroup:  input.NameGroup,
		MajorID:    input.MajorID,
		SemesterID: input.SemesterID,
		LeaderID:   currentStudent.ID,
		Status:     constant.CapstoneGroupStatus.ReviewingTopic,
	}

	if err := global.Db.Model(model.CapstoneGroup{}).Create(&data).Error; err != nil {
		return nil, err
	}

	studentCapstoneGroups := make([]model.StudentCapstoneGroup, 0)
	for _, id := range memberGroupsIDs {
		studentCapstoneGroups = append(studentCapstoneGroups, model.StudentCapstoneGroup{
			StudentID:       id,
			SemesterID:      input.SemesterID,
			CapstoneGroupID: data.ID,
		})
	}
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Create(&studentCapstoneGroups).Error; err != nil {
		return nil, err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", data.ID).First(&capstoneGroup).Error; err != nil {
		return nil, err
	}

	capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupOutput(&capstoneGroup)

	return &capstoneGroupOutput, nil
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
		NameGroup: input.NameGroup,
	}).Error; err != nil {
		return err
	}

	return nil
}

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

	token, err := jwt_util.GenerateInviteMentorToken(jwt_util.InviteMentorJwtInput{
		TeacherID:       input.TeacherID,
		CapstoneGroupID: input.CapstoneGroupID,
	})
	if err != nil {
		return err
	}

	// TODO: Create expired token and Rate Limit to prevent spam
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

func (cgs *capstoneGroupService) AcceptInviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.AcceptInviteMentorToCapstoneGroupInput) error {
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

	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).Updates(&model.CapstoneGroup{
		MentorID: &currentTeacher.ID,
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

func (cgs *capstoneGroupService) GetMentorAndListMemberCapstoneGroup(ctx *gin.Context, id int64) (*capstone_group_dto.MentorAndListMemberCapstoneGroupOutput, error) {
	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Joins("Mentor.User").Where("capstone_groups.id = ?", id).First(&capstoneGroup).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}
	mentor := capstoneGroup.Mentor
	var mentorOutput *user_dto.TeacherOutput
	mentorOutput = nil
	if mentor != nil {
		mentorOutput = user_dto.ToTeacherOutput(mentor)
	}

	var membersCapstoneGroup []model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Joins("Student.User").
		Where("student_capstone_groups.capstone_group_id = ?", id).
		Find(&membersCapstoneGroup).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	var output capstone_group_dto.MentorAndListMemberCapstoneGroupOutput
	output.LeaderID = capstoneGroup.LeaderID
	output.Mentor = mentorOutput
	output.Members = make([]*user_dto.StudentOutput, len(membersCapstoneGroup))

	for index, memberCapstoneGroup := range membersCapstoneGroup {
		member := memberCapstoneGroup.Student
		memberOutput := user_dto.ToStudentOutput(&member)
		output.Members[index] = memberOutput
	}

	return &output, nil
}
