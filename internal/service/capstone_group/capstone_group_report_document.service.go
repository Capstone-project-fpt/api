package capstone_group_service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

func (cgs *capstoneGroupService) CreateCapstoneGroupReportDocument(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupReportDocumentInput) error {
	err := cgs.validatePermissionCurrentStudentCapstoneGroup(ctx, input.CapstoneGroupID)
	if err != nil {
		return err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if capstoneGroup.Status != constant.CapstoneGroupStatus.InProgress {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupInReviewingTopicProcess,
		}))
	}

	var exitReportDocument model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("capstone_group_id = ? AND type_report = ?", input.CapstoneGroupID, input.TypeReport).First(&exitReportDocument).Error; err == nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.AlreadyExitDocumentReport,
		}))
	}

	reportDocument := model.ReportDocument{
		CapstoneGroupID:    input.CapstoneGroupID,
		TypeReport:         input.TypeReport,
		Name:               input.Name,
		FileIDs:            input.FileIDs,
		MentorReviewStatus: constant.MentorReviewStatusReport.Reviewing,
	}

	if err := global.Db.Model(model.ReportDocument{}).Create(&reportDocument).Error; err != nil {
		return err
	}

	var studentCapstoneGroups []model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Where("capstone_group_id = ?", input.CapstoneGroupID).Find(&studentCapstoneGroups).Error; err != nil {
		return err
	}

	var studentReportDocumentScores []model.ReportDocumentStudentScore

	for _, studentCapstoneGroup := range studentCapstoneGroups {
		studentReportDocumentScores = append(studentReportDocumentScores, model.ReportDocumentStudentScore{
			StudentID:        studentCapstoneGroup.StudentID,
			ReportDocumentID: reportDocument.ID,
			Score:            nil,
		})
	}

	if err := global.Db.Model(model.ReportDocumentStudentScore{}).Create(&studentReportDocumentScores).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) UpdateCapstoneGroupReportDocument(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupReportDocumentInput) error {
	err := cgs.validatePermissionCurrentStudentCapstoneGroup(ctx, input.CapstoneGroupID)
	if err != nil {
		return err
	}

	var reportDocument model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", input.ID).First(&reportDocument).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", input.ID).Updates(&model.ReportDocument{
		Name:    input.Name,
		FileIDs: input.FileIDs,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) DeleteCapstoneGroupReportDocument(ctx *gin.Context, reportID int64, capstoneGroupID int64) error {
	err := cgs.validatePermissionCurrentStudentCapstoneGroup(ctx, capstoneGroupID)
	if err != nil {
		return err
	}

	var reportDocument model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", reportID).First(&reportDocument).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", reportID).Delete(&model.ReportDocument{}).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) GetCapstoneGroupReportDocument(ctx *gin.Context, id int64) (*capstone_group_dto.ReportDocumentOutput, error) {
	var reportDocument model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", id).First(&reportDocument).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	return capstone_group_dto.ToReportDocumentOutput(&reportDocument), nil
}

func (cgs *capstoneGroupService) GetCapstoneGroupReportDocuments(ctx *gin.Context, capstoneGroupID int64) ([]capstone_group_dto.ReportDocumentOutput, error) {
	var reportDocuments []model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("capstone_group_id = ?", capstoneGroupID).Find(&reportDocuments).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	var reportDocumentOutputs []capstone_group_dto.ReportDocumentOutput
	for _, reportDocument := range reportDocuments {
		reportDocumentOutputs = append(reportDocumentOutputs, *capstone_group_dto.ToReportDocumentOutput(&reportDocument))
	}

	return reportDocumentOutputs, nil
}

func (cgs *capstoneGroupService) MentorUpdateStudentScoreForReportDocument(ctx *gin.Context, input *capstone_group_dto.MentorUpdateStudentScoreForReportDocument) error {
	currentTeacher, err := cgs.getCurrentTeacher(ctx)
	if err != nil {
		return err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.CapstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if *capstoneGroup.MentorID != currentTeacher.ID {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	var allStudentInCapstoneGroup []model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Where("capstone_group_id = ?", input.CapstoneGroupID).Find(&allStudentInCapstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if len(allStudentInCapstoneGroup) != len(input.StudentScoreData) {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
	}

	allStudentIDInCapstoneGroup := funk.Map(allStudentInCapstoneGroup, func(studentCapstoneGroup model.StudentCapstoneGroup) int64 {
		return studentCapstoneGroup.StudentID
	}).([]int64)

	for _, studentScoreData := range input.StudentScoreData {
		if studentScoreData.Score < 0 {
			return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InvalidParams,
			}))
		}

		if studentScoreData.Score > 10 {
			return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InvalidParams,
			}))
		}

		if !funk.Contains(allStudentIDInCapstoneGroup, studentScoreData.StudentID) {
			return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InvalidParams,
			}))
		}
	}

	var studentScores []model.ReportDocumentStudentScore
	if err := global.Db.Model(model.ReportDocumentStudentScore{}).Where("report_document_id = ?", input.ReportDocumentID).Find(&studentScores).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	for _, stundetScore := range studentScores {
		for _, studentScoreData := range input.StudentScoreData {
			if stundetScore.StudentID == studentScoreData.StudentID {
				stundetScore.Score = &studentScoreData.Score
				if err := global.Db.Save(&stundetScore).Error; err != nil {
					return err
				}
			}
		}
	}

	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", input.ReportDocumentID).Update("status", constant.MentorReviewStatusReport.Done).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) GetReportDocumentStudentsScore(ctx *gin.Context, reportDocumentID int64) ([]capstone_group_dto.ReportDocumentStudentScoreOutput, error) {
	var studentScores []model.ReportDocumentStudentScore
	query := global.Db.Model(model.ReportDocumentStudentScore{}).
		InnerJoins("Student.User").
		Where("report_document_id = ?", reportDocumentID).
		Find(&studentScores)

	if err := query.Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	var reportDocumentStudentScoreOutputs []capstone_group_dto.ReportDocumentStudentScoreOutput
	for _, studentScore := range studentScores {
		reportDocumentStudentScoreOutputs = append(reportDocumentStudentScoreOutputs, *capstone_group_dto.ToReportDocumentStudentScoreOutput(&studentScore))
	}

	return reportDocumentStudentScoreOutputs, nil
}

func (cgs *capstoneGroupService) AdminUpdateStudentScoreForReportDocument(ctx *gin.Context, input *capstone_group_dto.AdminUpdateStudentScoreForReportDocument) error {
	var studentScore model.ReportDocumentStudentScore
	if err := global.Db.Model(model.ReportDocumentStudentScore{}).Where("id = ?", input.ID).First(&studentScore).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	if input.Score < 0 {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
	}

	if input.Score > 10 {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
	}

	studentScore.Score = &input.Score
	if err := global.Db.Save(&studentScore).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) validatePermissionCurrentStudentCapstoneGroup(ctx *gin.Context, capstoneGroupID int64) error {
	currentStudent, err := cgs.getCurrentStudent(ctx)
	if err != nil {
		return err
	}

	var studentCapstoneGroup model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Where("student_id = ? AND capstone_group_id = ?", currentStudent.ID, capstoneGroupID).First(&studentCapstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	return nil
}
