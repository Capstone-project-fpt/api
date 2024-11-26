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
	GetListScheduleReview(ctx *gin.Context, input *schedule_review_dto.GetListScheduleReviewInput) (*[]schedule_review_dto.ScheduleReviewDetailOutput, error)
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

	tx := global.Db.Begin()
	if tx.Error != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return errors.New(message)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	internalMessageError := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.InternalServerError,
	})

	scheduleReview := model.ScheduleReview{
		Title:                 input.Title,
		Description:           input.Description,
		LinkMeeting:           linkMeeting,
		Type:                  input.Type,
		StartTime:             input.StartTime,
		EndTime:               input.EndTime,
		EvaluationCommitteeID: input.EvaluationCommitteeID,
		CapstoneGroupID:       input.CapstoneGroupID,
	}

	if err := tx.Model(model.ScheduleReview{}).Create(&scheduleReview).Error; err != nil {
		tx.Rollback()
		return errors.New(internalMessageError)
	}

	if err := tx.Model(model.CapstoneGroupReview{}).Create(&model.CapstoneGroupReview{
		CapstoneGroupID:  input.CapstoneGroupID,
		ScheduleReviewID: scheduleReview.ID,
	}).Error; err != nil {
		tx.Rollback()
		return errors.New(internalMessageError)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New(internalMessageError)
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

	var capstoneGroupReview model.CapstoneGroupReview
	queryCapstoneGroupReview := global.Db.Model(model.CapstoneGroupReview{}).
		Where("schedule_review_id = ?", scheduleReview.ID).
		First(&capstoneGroupReview)
	if err := queryCapstoneGroupReview.Error; err != nil {
		return nil, err
	}

	capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupOutput(&scheduleReview.CapstoneGroup)
	evaluationCommitteeOutput := evaluation_committee_dto.ToEvaluationCommitteeWithTeacherInfoOutput(&scheduleReview.EvaluationCommittee, &teacherOutput)
	capstoneGroupReviewOutput := capstone_group_dto.ToCapstoneGroupReviewOutput(&capstoneGroupReview)

	return schedule_review_dto.ToScheduleReviewDetailOutput(&scheduleReview, evaluationCommitteeOutput, &capstoneGroupOutput, capstoneGroupReviewOutput), nil
}

func (s *scheduleReviewService) GetListScheduleReview(ctx *gin.Context, input *schedule_review_dto.GetListScheduleReviewInput) (*[]schedule_review_dto.ScheduleReviewDetailOutput, error) {
	var scheduleReviews []model.ScheduleReview
	queryScheduleReviews := global.Db.Model(model.ScheduleReview{}).
		Select(`
			"schedule_reviews"."id",
  		"schedule_reviews"."title",
  		"schedule_reviews"."description",
  		"schedule_reviews"."link_meeting",
  		"schedule_reviews"."type",
  		"schedule_reviews"."start_time",
  		"schedule_reviews"."end_time",
  		"schedule_reviews"."evaluation_committee_id",
  		"schedule_reviews"."capstone_group_id",
  		"schedule_reviews"."created_at",
  		"schedule_reviews"."updated_at",
  		"CapstoneGroup"."id" AS "CapstoneGroup__id",
  		"CapstoneGroup"."name_group" AS "CapstoneGroup__name_group",
  		"CapstoneGroup"."topic_id" AS "CapstoneGroup__topic_id",
  		"CapstoneGroup"."major_id" AS "CapstoneGroup__major_id",
  		"CapstoneGroup"."semester_id" AS "CapstoneGroup__semester_id",
  		"CapstoneGroup"."leader_id" AS "CapstoneGroup__leader_id",
  		"CapstoneGroup"."mentor_id" AS "CapstoneGroup__mentor_id",
  		"CapstoneGroup"."status" AS "CapstoneGroup__status",
  		"CapstoneGroup"."created_at" AS "CapstoneGroup__created_at",
  		"CapstoneGroup"."updated_at" AS "CapstoneGroup__updated_at",
  		"EvaluationCommittee"."id" AS "EvaluationCommittee__id",
  		"EvaluationCommittee"."name" AS "EvaluationCommittee__name",
  		"EvaluationCommittee"."teacher_ids" AS "EvaluationCommittee__teacher_ids",
  		"EvaluationCommittee"."assign_group_ids" AS "EvaluationCommittee__assign_group_ids",
  		"EvaluationCommittee"."semester_id" AS "EvaluationCommittee__semester_id",
  		"EvaluationCommittee"."created_at" AS "EvaluationCommittee__created_at",
  		"EvaluationCommittee"."updated_at" AS "EvaluationCommittee__updated_at",
			"CapstoneGroupReview"."id" AS "CapstoneGroupReview__ID",
			"CapstoneGroupReview"."capstone_group_id" AS "CapstoneGroupReview__CapstoneGroupID",
			"CapstoneGroupReview"."schedule_review_id" AS "CapstoneGroupReview__ScheduleReviewID",
			"CapstoneGroupReview"."report_files" AS "CapstoneGroupReview__ReportFiles",
			"CapstoneGroupReview"."feedback" AS "CapstoneGroupReview__Feedback",
			"CapstoneGroupReview"."created_at" AS "CapstoneGroupReview__CreatedAt",
			"CapstoneGroupReview"."updated_at" AS "CapstoneGroupReview__UpdatedAt"
		`).
		InnerJoins(`INNER JOIN "capstone_group_reviews" as "CapstoneGroupReview" ON "schedule_reviews"."id" = "CapstoneGroupReview"."schedule_review_id"`).
		InnerJoins("CapstoneGroup").
		InnerJoins("EvaluationCommittee").
		Where("start_time >= ? AND end_time <= ?", input.StartTime, input.EndTime)

	if input.CapstoneGroupID != nil {
		queryScheduleReviews = queryScheduleReviews.Where("schedule_reviews.capstone_group_id = ?", *input.CapstoneGroupID)
	}

	if input.EvaluationCommitteeID != nil {
		queryScheduleReviews = queryScheduleReviews.Where("schedule_reviews.evaluation_committee_id = ?", *input.EvaluationCommitteeID)
	}

	if input.OrderBy != nil {
		queryScheduleReviews.Order("start_time " + *input.OrderBy)
	}

	if err := queryScheduleReviews.Find(&scheduleReviews).Error; err != nil {
		return nil, err
	}

	evaluationCommittees := make(map[int64]model.EvaluationCommittee)
	for i := 0; i < len(scheduleReviews); i++ {
		evaluationCommittees[scheduleReviews[i].EvaluationCommitteeID] = scheduleReviews[i].EvaluationCommittee
	}

	evaluationCommitteeWithTeacherInfoOutput := make(map[int64]*evaluation_committee_dto.EvaluationCommitteeWithTeacherInfoOutput)

	for evaluationCommitteeID, evaluationCommittee := range evaluationCommittees {
		var teachers []model.Teacher
		queryTeachers := global.Db.Model(model.Teacher{}).
			Joins("User").
			Where("teachers.id IN ?", []int64(evaluationCommittee.TeacherIDs)).
			Find(&teachers)
		if err := queryTeachers.Error; err != nil {
			return nil, err
		}

		var teacherOutput []*user_dto.TeacherOutput
		for _, teacher := range teachers {
			teacherOutput = append(teacherOutput, user_dto.ToTeacherOutput(&teacher))
		}

		evaluationCommitteeWithTeacherInfoOutput[evaluationCommitteeID] = evaluation_committee_dto.ToEvaluationCommitteeWithTeacherInfoOutput(&evaluationCommittee, &teacherOutput)
	}

	var scheduleReviewOutputs []schedule_review_dto.ScheduleReviewDetailOutput
	for _, scheduleReview := range scheduleReviews {
		capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupOutput(&scheduleReview.CapstoneGroup)
		capstoneGroupReviewOutput := capstone_group_dto.ToCapstoneGroupReviewOutput(scheduleReview.CapstoneGroupReview)

		scheduleReviewOutputs = append(
			scheduleReviewOutputs,
			*schedule_review_dto.ToScheduleReviewDetailOutput(
				&scheduleReview,
				evaluationCommitteeWithTeacherInfoOutput[scheduleReview.EvaluationCommitteeID],
				&capstoneGroupOutput,
				capstoneGroupReviewOutput,
			),
		)
	}

	return &scheduleReviewOutputs, nil
}
