package capstone_group_service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (cgs *capstoneGroupService) CreateCapstoneGroupReportDocument(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupReportDocumentInput) error {
	err := cgs.validatePermissionCurrentStudentCapstoneGroup(ctx, input.CapstoneGroupID)
	if err != nil {
		return err
	}

	var exitReportDocument model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("capstone_group_id = ? AND type_report = ?", input.CapstoneGroupID, input.TypeReport).First(&exitReportDocument).Error; err == nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.AlreadyExitDocumentReport,
		}))
	}

	if err := global.Db.Model(model.ReportDocument{}).Create(&model.ReportDocument{
		CapstoneGroupID:    input.CapstoneGroupID,
		TypeReport:         input.TypeReport,
		Name:               input.Name,
		FileIDs:            input.FileIDs,
		MentorReviewStatus: constant.MentorReviewStatusReport.Reviewing,
	}).Error; err != nil {
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
