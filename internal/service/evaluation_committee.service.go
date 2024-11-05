package service

import (
	"errors"
	"strings"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/evaluation_committee_dto"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

type IEvaluationCommitteeService interface {
	CreateEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.CreateEvaluationCommitteeInput) error
	GetEvaluationCommittee(ctx *gin.Context, id int64) (*evaluation_committee_dto.EvaluationCommitteeOutput, error)
}

type evaluationCommitteeService struct{}

func NewEvaluationCommitteeService() IEvaluationCommitteeService {
	return &evaluationCommitteeService{}
}

func (e *evaluationCommitteeService) CreateEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.CreateEvaluationCommitteeInput) error {
	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", input.SemesterID).First(&semester).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		})

		return errors.New(message)
	}

	var teachers []model.Teacher
	queryTeachers := global.Db.Model(model.Teacher{}).Joins("User").Where("teachers.id IN ?", input.TeacherIDs).Find(&teachers)
	if err := queryTeachers.Error; err != nil {
		return err
	}

	if len(teachers) != len(input.TeacherIDs) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return errors.New(message)
	}

	var allEvaluationCommittees []model.EvaluationCommittee
	queryEvaluationCommittees := global.Db.Model(model.EvaluationCommittee{}).Where("semester_id = ?", input.SemesterID).Find(&allEvaluationCommittees)
	if err := queryEvaluationCommittees.Error; err != nil {
		return err
	}

	assignedTeacherIDs := make([]int64, 0)
	totalAssignedTeacher := 0
	for _, evaluationCommittee := range allEvaluationCommittees {
		totalAssignedTeacher += len(evaluationCommittee.TeacherIDs)
		assignedTeacherIDs = append(assignedTeacherIDs, evaluationCommittee.TeacherIDs...)
	}

	teachersHadAssigned := make([]model.Teacher, 0)

	for _, teacher := range teachers {
		isAssigned := funk.Contains(assignedTeacherIDs, teacher.ID)

		if isAssigned {
			teachersHadAssigned = append(teachersHadAssigned, teacher)
		}
	}

	if len(teachersHadAssigned) > 0 {
		teacherNames := funk.Map(teachersHadAssigned, func(teacher model.Teacher) string {
			return teacher.User.Name
		}).([]string)

		teacherNamesStr := strings.Join(teacherNames, ", ")

		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.EvaluationCommitteeTeacherHadAssigned,
			TemplateData: map[string]interface{}{
				"TeacherNames": teacherNamesStr,
			},
		})
		return errors.New(message)
	}

	evaluationCommittee := model.EvaluationCommittee{
		Name:       input.Name,
		SemesterID: input.SemesterID,
		TeacherIDs: input.TeacherIDs,
	}

	if err := global.Db.Model(model.EvaluationCommittee{}).Create(&evaluationCommittee).Error; err != nil {
		return err
	}

	return nil
}

func (e *evaluationCommitteeService) GetEvaluationCommittee(ctx *gin.Context, id int64) (*evaluation_committee_dto.EvaluationCommitteeOutput, error) {
	var evaluationCommittee model.EvaluationCommittee

	query := global.Db.Model(&model.EvaluationCommittee{}).Where("id = ?", id).First(&evaluationCommittee)

	if err := query.Error; err != nil {
		return nil, err
	}

	output := evaluation_committee_dto.ToEvaluationCommitteeOutput(&evaluationCommittee)

	return output, nil
}
