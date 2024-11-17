package service

import (
	"errors"
	"strings"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/evaluation_committee_dto"
	"github.com/api/internal/dto/user_dto"
	util "github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

type IEvaluationCommitteeService interface {
	CreateEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.CreateEvaluationCommitteeInput) error
	UpdateEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.UpdateEvaluationCommitteeInput) error
	DeleteEvaluationCommittee(ctx *gin.Context, id int64) error
	GetEvaluationCommittee(ctx *gin.Context, id int64) (*evaluation_committee_dto.EvaluationCommitteeWithTeacherInfoOutput, error)
	GetListEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.GetListEvaluationCommitteeInput) (*evaluation_committee_dto.ListEvaluationCommitteeOutput, error)
	GetListTeachersHaveEvaluationCommitteeGroup(ctx *gin.Context, semesterID int64) (*[]*user_dto.TeacherOutput, error)
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
	// teacherIDs :=
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

func (e *evaluationCommitteeService) UpdateEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.UpdateEvaluationCommitteeInput) error {
	var evaluationCommittee model.EvaluationCommittee
	if err := global.Db.Model(model.EvaluationCommittee{}).Where("id = ?", input.ID).First(&evaluationCommittee).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.EvaluationCommitteeNotFound,
		})
		return errors.New(message)
	}

	if !util.IsSameElementTwoArray(input.TeacherIDs, evaluationCommittee.TeacherIDs) {
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

		var allEvaluationCommitteesInSemester []model.EvaluationCommittee

		if err := global.Db.Model(model.EvaluationCommittee{}).Where("semester_id = ? AND id != ?", evaluationCommittee.SemesterID, evaluationCommittee.ID).Find(&allEvaluationCommitteesInSemester).Error; err != nil {
			return err
		}

		allTeachersInEvaluationCommittee := make([]int64, 0)
		for _, evaluationCommittee := range allEvaluationCommitteesInSemester {
			allTeachersInEvaluationCommittee = append(allTeachersInEvaluationCommittee, evaluationCommittee.TeacherIDs...)
		}

		assignedTeacherNames := make([]string, 0)
		for _, teacher := range teachers {
			if funk.Contains(allTeachersInEvaluationCommittee, teacher.ID) {
				assignedTeacherNames = append(assignedTeacherNames, teacher.User.Name)
			}
		}

		if len(assignedTeacherNames) > 0 {
			teacherNamesStr := strings.Join(assignedTeacherNames, ", ")
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.EvaluationCommitteeTeacherHadAssigned,
				TemplateData: map[string]interface{}{
					"TeacherNames": teacherNamesStr,
				},
			})
			return errors.New(message)
		}
	}

	evaluationCommittee.Name = input.Name
	evaluationCommittee.TeacherIDs = input.TeacherIDs

	if err := global.Db.Model(model.EvaluationCommittee{}).Where("id = ?", input.ID).Save(&evaluationCommittee).Error; err != nil {
		return err
	}

	return nil
}

func (e *evaluationCommitteeService) DeleteEvaluationCommittee(ctx *gin.Context, id int64) error {
	var evaluationCommittee model.EvaluationCommittee
	if err := global.Db.Model(model.EvaluationCommittee{}).Where("id = ?", id).First(&evaluationCommittee).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.EvaluationCommitteeNotFound,
		})
		return errors.New(message)
	}

	if err := global.Db.Model(model.EvaluationCommittee{}).Where("id = ?", id).Delete(&evaluationCommittee).Error; err != nil {
		return err
	}

	return nil
}

func (e *evaluationCommitteeService) GetEvaluationCommittee(ctx *gin.Context, id int64) (*evaluation_committee_dto.EvaluationCommitteeWithTeacherInfoOutput, error) {
	var evaluationCommittee model.EvaluationCommittee

	queryEvaluationCommittee := global.Db.Model(&model.EvaluationCommittee{}).Where("id = ?", id).First(&evaluationCommittee)

	if err := queryEvaluationCommittee.Error; err != nil {
		return nil, err
	}

	var teachers []model.Teacher

	queryTeachers := global.Db.Model(model.Teacher{}).
		Select(`
			"teachers"."id",
			"teachers"."sub_major_id",
			"teachers"."user_id",
			"User"."id" AS "User__id", 
  		"User"."name" AS "User__name", 
  		"User"."user_type" AS "User__user_type", 
  		"User"."email" AS "User__email", 
  		"User"."phone_number" AS "User__phone_number"
		`).
		Joins(`INNER JOIN "users" "User" ON "User"."id" = "teachers"."user_id"`).
		Where("teachers.id IN ?", []int64(evaluationCommittee.TeacherIDs)).
		Find(&teachers)

	if err := queryTeachers.Error; err != nil {
		return nil, err
	}

	var teacherOutput []*user_dto.TeacherOutput
	for _, teacher := range teachers {
		teacherOutput = append(teacherOutput, user_dto.ToTeacherOutput(&teacher))
	}

	output := evaluation_committee_dto.ToEvaluationCommitteeWithTeacherInfoOutput(&evaluationCommittee, &teacherOutput)

	return output, nil
}

func (e *evaluationCommitteeService) GetListEvaluationCommittee(ctx *gin.Context, input *evaluation_committee_dto.GetListEvaluationCommitteeInput) (*evaluation_committee_dto.ListEvaluationCommitteeOutput, error) {
	var total int64
	var items []model.EvaluationCommittee

	query := global.Db.Model(model.EvaluationCommittee{})
	queryTotal := global.Db.Model(model.EvaluationCommittee{})
	if input.SemesterID != nil {
		query.Where("semester_id = ?", *input.SemesterID)
		queryTotal.Where("semester_id = ?", *input.SemesterID)
	}

	if input.OrderBy != nil {
		query.Order("evaluation_committees.created_at " + *input.OrderBy)
	}

	if err := queryTotal.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := query.Limit(int(input.Limit)).Offset(int(input.Offset)).Find(&items).Error; err != nil {
		return nil, err
	}

	itemsEvaluationCommitteeOutput := make([]*evaluation_committee_dto.EvaluationCommitteeOutput, len(items))
	for i, item := range items {
		itemsEvaluationCommitteeOutput[i] = evaluation_committee_dto.ToEvaluationCommitteeOutput(&item)
	}

	return &evaluation_committee_dto.ListEvaluationCommitteeOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsEvaluationCommitteeOutput,
	}, nil
}

func (e *evaluationCommitteeService) GetListTeachersHaveEvaluationCommitteeGroup(ctx *gin.Context, semesterID int64) (*[]*user_dto.TeacherOutput, error) {
	var evaluationCommittees []model.EvaluationCommittee
	queryEvaluationCommittees := global.Db.Model(&model.EvaluationCommittee{}).Where("semester_id = ?", semesterID).First(&evaluationCommittees)
	if err := queryEvaluationCommittees.Error; err != nil {
		return nil, err
	}

	teacherIDs := make([]int64, 0)
	for _, evaluationCommittee := range evaluationCommittees {
		for _, teacherID := range evaluationCommittee.TeacherIDs {
			if !funk.Contains(teacherIDs, teacherID) {
				teacherIDs = append(teacherIDs, teacherID)
			}
		}
	}

	var teachers []model.Teacher

	queryTeachers := global.Db.Model(model.Teacher{}).
		Select(`
			"teachers"."id",
			"teachers"."sub_major_id",
			"teachers"."user_id",
			"User"."id" AS "User__id", 
  			"User"."name" AS "User__name", 
  			"User"."user_type" AS "User__user_type", 
  			"User"."email" AS "User__email", 
  			"User"."phone_number" AS "User__phone_number"
		`).
		Joins(`INNER JOIN "users" "User" ON "User"."id" = "teachers"."user_id"`).
		Where("teachers.id IN ?", []int64(teacherIDs)).
		Find(&teachers)

	if err := queryTeachers.Error; err != nil {
		return nil, err
	}

	var teachersHaveEvaluationCommitteeGroup []*user_dto.TeacherOutput
	for _, teacher := range teachers {
		teachersHaveEvaluationCommitteeGroup = append(teachersHaveEvaluationCommitteeGroup, user_dto.ToTeacherOutput(&teacher))
	}
	return &teachersHaveEvaluationCommitteeGroup, nil

}
