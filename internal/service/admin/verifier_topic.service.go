package service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/internal/dto/user_dto"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (as *adminService) AssignVerifierTopic(ctx *gin.Context, input *admin_dto.AdminAssignVerifierTopicInput) error {
	var teacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Select("id").Where("id = ?", input.TeacherID).First(&teacher).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return errors.New(message)
	}

	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Select("id").Where("id = ?", input.SemesterID).First(&semester).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		})

		return errors.New(message)
	}

	var exitVerifierTopic model.VerifierTopic
	if err := global.Db.Model(model.VerifierTopic{}).Select("id").Where("teacher_id = ? AND semester_id = ?", input.TeacherID, input.SemesterID).First(&exitVerifierTopic).Error; err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.VerifierTopicAlreadyExist,
		})

		return errors.New(message)
	}

	verifierTopic := model.VerifierTopic{
		TeacherID:  input.TeacherID,
		SemesterID: input.SemesterID,
	}

	if err := global.Db.Create(&verifierTopic).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return errors.New(message)
	}

	return nil
}

func (as *adminService) UnassignVerifierTopic(ctx *gin.Context, input *admin_dto.AdminUnassignVerifierTopicInput) error {
	var verifierTopic model.VerifierTopic
	if err := global.Db.Model(model.VerifierTopic{}).Select("id").Where("teacher_id = ? AND semester_id = ?", input.TeacherID, input.SemesterID).First(&verifierTopic).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.VerifierTopicNotFound,
		})

		return errors.New(message)
	}

	if err := global.Db.Delete(&verifierTopic).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return errors.New(message)
	}

	return nil
}

func (as *adminService) GetListVerifiersTopic(ctx *gin.Context, semesterID int64) (*[]*user_dto.TeacherOutput, error) {
	var verifierTopics []model.VerifierTopic
	queryVerifierTopics := global.Db.Model(model.VerifierTopic{}).
		InnerJoins("Teacher.User").
		Where("semester_id = ?", semesterID).
		Find(&verifierTopics)

	if err := queryVerifierTopics.Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return nil, errors.New(message)
	}

	teachersOutput := make([]*user_dto.TeacherOutput, len(verifierTopics))
	for index, verifierTopic := range verifierTopics {
		teachersOutput[index] = user_dto.ToTeacherOutput(&verifierTopic.Teacher)
	}

	return &teachersOutput, nil
}
