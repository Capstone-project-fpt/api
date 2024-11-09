package service

import (
	"errors"
	"time"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/internal/dto/evaluation_committee_dto"
	"github.com/api/internal/dto/schedule_review_dto"
	"github.com/api/internal/dto/user_dto"
	"github.com/api/pkg/service/google"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type IScheduleReviewService interface {
	CreateScheduleReview(ctx *gin.Context, input *schedule_review_dto.CreateScheduleReviewInput) error
	UpdateScheduleReview(ctx *gin.Context, input *schedule_review_dto.UpdateScheduleReviewInput) error
	DeleteScheduleReview(ctx *gin.Context, id int64) error
	GetScheduleReviewDetail(ctx *gin.Context, id int64) (*schedule_review_dto.ScheduleReviewDetailOutput, error)
	GetListScheduleReview(ctx *gin.Context, input *schedule_review_dto.GetListScheduleReviewInput) (*[]schedule_review_dto.ScheduleReviewOutput, error)
}

type scheduleReviewService struct {
	googleService google.IGoogleService
}

func NewScheduleReviewService(
	googleService google.IGoogleService,
) IScheduleReviewService {
	return &scheduleReviewService{
		googleService: googleService,
	}
}

type ValidateScheduleReview struct {
	ScheduleReviewID      *int64
	Type                  string
	StartTime             time.Time
	EndTime               time.Time
	EvaluationCommitteeID int64
	CapstoneGroupID       int64
	SemesterID            int64
}

func (s *scheduleReviewService) CreateScheduleReview(ctx *gin.Context, input *schedule_review_dto.CreateScheduleReviewInput) error {
	// TODO: Create redis key to check create calendar in progress

	inputValidate := &ValidateScheduleReview{
		ScheduleReviewID:      nil,
		Type:                  input.Type,
		StartTime:             input.StartTime,
		EndTime:               input.EndTime,
		EvaluationCommitteeID: input.EvaluationCommitteeID,
		CapstoneGroupID:       input.CapstoneGroupID,
		SemesterID:            input.SemesterID,
	}

	if err := s.validateScheduleReviewBeforeUpsert(inputValidate); err != nil {
		return err
	}

	linkMeeting := s.googleService.GenerateGoogleMeeting()

	if err := global.Db.Model(model.ScheduleReview{}).Create(&model.ScheduleReview{
		Title:                 input.Title,
		Description:           input.Description,
		LinkMeeting:           linkMeeting,
		Type:                  input.Type,
		StartTime:             input.StartTime,
		EndTime:               input.EndTime,
		EvaluationCommitteeID: input.EvaluationCommitteeID,
		CapstoneGroupID:       input.CapstoneGroupID,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *scheduleReviewService) UpdateScheduleReview(ctx *gin.Context, input *schedule_review_dto.UpdateScheduleReviewInput) error {
	// TODO: Create redis key to check create calendar in progress

	inputValidate := &ValidateScheduleReview{
		ScheduleReviewID:      &input.ID,
		Type:                  input.Type,
		StartTime:             input.StartTime,
		EndTime:               input.EndTime,
		EvaluationCommitteeID: input.EvaluationCommitteeID,
		CapstoneGroupID:       input.CapstoneGroupID,
		SemesterID:            input.SemesterID,
	}

	if err := s.validateScheduleReviewBeforeUpsert(inputValidate); err != nil {
		return err
	}

	linkMeeting := s.googleService.GenerateGoogleMeeting()

	if err := global.Db.Model(model.ScheduleReview{}).Where("id = ?", input.ID).Updates(&model.ScheduleReview{
		Title:                 input.Title,
		Description:           input.Description,
		Type:                  input.Type,
		LinkMeeting:           linkMeeting,
		StartTime:             input.StartTime,
		EndTime:               input.EndTime,
		EvaluationCommitteeID: input.EvaluationCommitteeID,
		CapstoneGroupID:       input.CapstoneGroupID,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *scheduleReviewService) validateScheduleReviewBeforeUpsert(input *ValidateScheduleReview) error {
	var capstoneGroup model.CapstoneGroup
	var evaluationCommittee model.EvaluationCommittee
	var semester model.Semester

	if input.StartTime.After(input.EndTime) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ScheduleReviewStartTimeNeedToBeBeforeSemesterEndTime,
		})
		return errors.New(message)
	}

	querySemester := global.Db.Model(model.Semester{}).
		Where("id = ?", input.SemesterID).
		First(&semester)
	if err := querySemester.Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		})
		return errors.New(message)
	}

	if input.StartTime.Before(semester.StartTime) || input.EndTime.After(semester.EndTime) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ScheduleReviewTimeNeedToBeInSemesterTime,
		})
		return errors.New(message)
	}

	queryCapstoneGroup := global.Db.Model(model.CapstoneGroup{}).
		Select("id").
		Where("id = ?", input.CapstoneGroupID).
		Where("semester_id = ?", input.SemesterID).
		First(&capstoneGroup)
	if err := queryCapstoneGroup.Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		})
		return errors.New(message)
	}

	queryEvaluationCommittee := global.Db.Model(model.EvaluationCommittee{}).
		Select("id").
		Where("id = ?", input.EvaluationCommitteeID).
		Where("semester_id = ?", input.SemesterID).
		First(&evaluationCommittee)
	if err := queryEvaluationCommittee.Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.EvaluationCommitteeNotFound,
		})
		return errors.New(message)
	}

	var exitScheduleReviewTypeCapstoneGroup model.ScheduleReview
	queryScheduleReviewTypeCapstoneGroup := global.Db.Model(model.ScheduleReview{}).
		Select("id").
		Where("type = ?", input.Type).
		Where("capstone_group_id = ?", input.CapstoneGroupID)
	if input.ScheduleReviewID != nil {
		queryScheduleReviewTypeCapstoneGroup.Where("id != ?", *input.ScheduleReviewID)
	}
	if err := queryScheduleReviewTypeCapstoneGroup.First(&exitScheduleReviewTypeCapstoneGroup).Error; err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupAlreadyHadScheduleReview,
		})

		return errors.New(message)
	}

	var overlapScheduleReviewForCapstoneGroup model.ScheduleReview
	queryOverlapScheduleReviewForCapstoneGroup := global.Db.Model(model.ScheduleReview{}).
		Select("id").
		Where("start_time < ? AND end_time > ? AND capstone_group_id = ?", input.EndTime, input.StartTime, input.CapstoneGroupID)
	if input.ScheduleReviewID != nil {
		queryOverlapScheduleReviewForCapstoneGroup.Where("id != ?", *input.ScheduleReviewID)
	}
	if err := queryOverlapScheduleReviewForCapstoneGroup.First(&overlapScheduleReviewForCapstoneGroup).Error; err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ScheduleReviewCapstoneGroupOverlap,
		})

		return errors.New(message)
	}

	var overlapScheduleReviewForEvaluationCommittee model.ScheduleReview
	queryOverlapScheduleReviewForEvaluationCommittee := global.Db.Model(model.ScheduleReview{}).
		Select("id").
		Where("start_time < ? AND end_time > ? AND evaluation_committee_id = ?", input.EndTime, input.StartTime, input.EvaluationCommitteeID)
	if input.ScheduleReviewID != nil {
		queryOverlapScheduleReviewForEvaluationCommittee.Where("id != ?", *input.ScheduleReviewID)
	}
	if err := queryOverlapScheduleReviewForEvaluationCommittee.First(&overlapScheduleReviewForEvaluationCommittee).Error; err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ScheduleReviewEvaluationCommitteeOverlap,
		})

		return errors.New(message)
	}

	return nil
}

func (s *scheduleReviewService) DeleteScheduleReview(ctx *gin.Context, id int64) error {
	if err := global.Db.Model(model.ScheduleReview{}).Select("id").Where("id = ?", id).First(&model.ScheduleReview{}).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ScheduleReviewNotFound,
		})

		return errors.New(message)
	}

	if err := global.Db.Model(model.ScheduleReview{}).Where("id = ?", id).Delete(&model.ScheduleReview{}).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return errors.New(message)
	}

	return nil
}

func (s *scheduleReviewService) GetScheduleReviewDetail(ctx *gin.Context, id int64) (*schedule_review_dto.ScheduleReviewDetailOutput, error) {
	var scheduleReview model.ScheduleReview
	queryScheduleReview := global.Db.Model(model.ScheduleReview{}).
		Joins("CapstoneGroup").
		Joins("EvaluationCommittee").
		Where("schedule_reviews.id = ?", id).
		First(&scheduleReview)
	if err := queryScheduleReview.Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ScheduleReviewNotFound,
		})

		return nil, errors.New(message)
	}

	var teachers []model.Teacher
	queryTeachers := global.Db.Model(model.Teacher{}).
		Joins("User").
		Where("teachers.id IN ?", []int64(scheduleReview.EvaluationCommittee.TeacherIDs)).
		Find(&teachers)
	if err := queryTeachers.Error; err != nil {
		return nil, err
	}

	var teacherOutput []*user_dto.TeacherOutput
	for _, teacher := range teachers {
		teacherOutput = append(teacherOutput, user_dto.ToTeacherOutput(&teacher))
	}

	capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupOutput(&scheduleReview.CapstoneGroup)
	evaluationCommitteeOutput := evaluation_committee_dto.ToEvaluationCommitteeWithTeacherInfoOutput(&scheduleReview.EvaluationCommittee, &teacherOutput)

	return schedule_review_dto.ToScheduleReviewDetailOutput(&scheduleReview, evaluationCommitteeOutput, &capstoneGroupOutput), nil
}

func (s *scheduleReviewService) GetListScheduleReview(ctx *gin.Context, input *schedule_review_dto.GetListScheduleReviewInput) (*[]schedule_review_dto.ScheduleReviewOutput, error) {
	var scheduleReviews []model.ScheduleReview
	queryScheduleReviews := global.Db.Model(model.ScheduleReview{}).
		Where("start_time >= ? AND end_time <= ?", input.StartTime, input.EndTime)
	
	if input.CapstoneGroupID != nil {
		queryScheduleReviews = queryScheduleReviews.Where("capstone_group_id = ?", *input.CapstoneGroupID)
	}

	if input.EvaluationCommitteeID != nil {
		queryScheduleReviews = queryScheduleReviews.Where("evaluation_committee_id = ?", *input.EvaluationCommitteeID)
	}

	if input.OrderBy != nil {
		queryScheduleReviews.Order("start_time " + *input.OrderBy)
	}

	if err := queryScheduleReviews.Find(&scheduleReviews).Error; err != nil {
		return nil, err
	}

	var scheduleReviewOutputs []schedule_review_dto.ScheduleReviewOutput
	for _, scheduleReview := range scheduleReviews {
		scheduleReviewOutputs = append(scheduleReviewOutputs, *schedule_review_dto.ToScheduleReviewOutput(&scheduleReview))
	}

	return &scheduleReviewOutputs, nil
}