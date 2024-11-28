package capstone_group_service

import (
	"errors"
	"time"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/capstone_group_dto"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (cgs *capstoneGroupService) CommentReport(ctx *gin.Context, input *capstone_group_dto.CommentReportInput) error {
	userID, _, _, err := cgs.validatePermissionActionInReportDocumentAndReturnUserInfo(ctx, input.CapstoneGroupID)
	if err != nil {
		return err
	}

	var reportDocument model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).Where("id = ?", input.ReportDocumentID).First(&reportDocument).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportDocumentNotFound,
		}))
	}

	now := time.Now().Unix()
	reportComment := model.ReportComment{
		ReportDocumentID: reportDocument.ID,
		UserID:           userID,
		Message:          input.Message,
		GroupComment:     now,
	}

	if err := global.Db.Create(&reportComment).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) UpdateCommentReport(ctx *gin.Context, input *capstone_group_dto.UpdateCommentReportInput) error {
	userID, _, _, err := cgs.validatePermissionActionInReportDocumentAndReturnUserInfo(ctx, input.CapstoneGroupID)
	if err != nil {
		return err
	}

	var reportComment model.ReportComment
	queryReportComment := global.Db.Model(model.ReportComment{}).
		Where("id = ? AND user_id = ? AND report_document_id = ?", input.ID, userID, input.ReportDocumentID)

	if err := queryReportComment.First(&reportComment).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportCommentNotFound,
		}))
	}

	if err := global.Db.Model(model.ReportComment{}).Where("id = ?", input.ID).Update("message", input.Message).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) DeleteCommentReport(ctx *gin.Context, input *capstone_group_dto.DeleteCommentReportInput) error {
	userID, _, _, err := cgs.validatePermissionActionInReportDocumentAndReturnUserInfo(ctx, input.CapstoneGroupID)
	if err != nil {
		return err
	}

	var reportComment model.ReportComment
	queryReportComment := global.Db.Model(model.ReportComment{}).
		Where("id = ? AND user_id = ? AND report_document_id = ?", input.ID, userID, input.ReportDocumentID)

	if err := queryReportComment.First(&reportComment).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.ReportCommentNotFound,
		}))
	}

	if err := queryReportComment.Delete(&reportComment).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) GetListReportComments(ctx *gin.Context, input *capstone_group_dto.GetListCommentReportInput) (*[]*capstone_group_dto.ReportCommentWithUserInfoOutput, error) {
	var reportComments []model.ReportComment

	var orderBy string
	if input.OrderBy == nil {
		orderBy = "DESC"
	} else {
		orderBy = *input.OrderBy
	}

	queryReportComments := global.Db.Model(model.ReportComment{}).
		InnerJoins("User").
		Where("report_document_id = ?", input.ReportDocumentID).
		Order("created_at " + orderBy)

	if err := queryReportComments.Find(&reportComments).Error; err != nil {
		return nil, err
	}

	reportCommentsOutput := make([]*capstone_group_dto.ReportCommentWithUserInfoOutput, len(reportComments))
	for i, reportComment := range reportComments {
		reportCommentsOutput[i] = capstone_group_dto.ToReportCommentWithUserInfoOutput(&reportComment)
	}

	return &reportCommentsOutput, nil
}

func (cgs *capstoneGroupService) validatePermissionActionInReportDocumentAndReturnUserInfo(ctx *gin.Context, capstoneGroupID int64) (int64, *model.Student, *model.Teacher, error) {
	currentStudent, currentTeacher, err := cgs.getCurrentStudentOrTeacher(ctx)

	if err != nil {
		return 0, nil, nil, err
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", capstoneGroupID).First(&capstoneGroup).Error; err != nil {
		return 0, nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if currentTeacher != nil && *capstoneGroup.MentorID != currentTeacher.ID {
		return 0, nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	if currentStudent != nil {
		var studentCapstoneGroups []model.StudentCapstoneGroup
		if err := global.Db.Model(model.StudentCapstoneGroup{}).Where("capstone_group_id = ?", capstoneGroupID).Find(&studentCapstoneGroups).Error; err != nil {
			return 0, nil, nil, err
		}

		for _, studentCapstoneGroup := range studentCapstoneGroups {
			if studentCapstoneGroup.SemesterID != currentStudent.ID {
				return 0, nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: constant.MessageI18nId.PermissionDenied,
				}))
			}
		}
	}

	var userID int64
	if currentTeacher != nil {
		userID = currentTeacher.UserID
	} else {
		userID = currentStudent.UserID
	}

	return userID, currentStudent, currentTeacher, nil
}

func (cgs *capstoneGroupService) getCurrentStudentOrTeacher(ctx *gin.Context) (*model.Student, *model.Teacher, error) {
	var currentTeacher model.Teacher
	var currentStudent model.Student

	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	if currentUser.UserType == constant.UserType.Teacher {
		if err := global.Db.Model(model.Teacher{}).Where("user_id = ?", currentUser.ID).First(&currentTeacher).Error; err != nil {
			return nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.UserNotFound,
			}))
		}

		currentTeacher.User = model.User{
			ID:       currentUser.ID,
			Email:    currentUser.Email,
			Name:     currentUser.Name,
			UserType: currentUser.UserType,
		}

		return nil, &currentTeacher, nil
	}

	if currentUser.UserType == constant.UserType.Student {
		if err := global.Db.Model(model.Student{}).Where("user_id = ?", currentUser.ID).First(&currentStudent).Error; err != nil {
			return nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.UserNotFound,
			}))
		}

		currentStudent.User = model.User{
			ID:       currentUser.ID,
			Email:    currentUser.Email,
			Name:     currentUser.Name,
			UserType: currentUser.UserType,
		}

		return &currentStudent, nil, nil
	}

	return nil, nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UserNotFound,
	}))
}
