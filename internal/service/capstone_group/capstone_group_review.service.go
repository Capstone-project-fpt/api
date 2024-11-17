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

type ICapstoneGroupReviewService interface {
	UpdateReportFilesCapstoneGroupReview(ctx *gin.Context, input *capstone_group_dto.UpdateReportFilesCapstoneGroupReviewInput) error
	FeedbackCapstoneGroupReview(ctx *gin.Context, input *capstone_group_dto.FeedbackCapstoneGroupReviewInput) error
	GetCapstoneGroupReview(ctx *gin.Context, id int64, capstoneGroupID int64) (*capstone_group_dto.CapstoneGroupReviewOutput, error)
}

type capstoneGroupReviewService struct{}

func NewCapstoneGroupReviewService() ICapstoneGroupReviewService {
	return &capstoneGroupReviewService{}
}

func (c *capstoneGroupReviewService) UpdateReportFilesCapstoneGroupReview(ctx *gin.Context, input *capstone_group_dto.UpdateReportFilesCapstoneGroupReviewInput) error {
	var capstoneGroupReview model.CapstoneGroupReview

	if err := global.Db.Model(model.CapstoneGroupReview{}).
		Where("id = ? AND capstone_group_id = ?", input.CapstoneGroupReviewID, input.CapstoneGroupID).
		First(&capstoneGroupReview).
		Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupReviewNotFound,
		}))
	}

	if err := global.Db.Model(model.CapstoneGroupReview{}).Where("id = ?", input.CapstoneGroupReviewID).Updates(&model.CapstoneGroupReview{
		ReportFiles: input.ReportFiles,
	}).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		}))
	}

	return nil
}

func (c *capstoneGroupReviewService) FeedbackCapstoneGroupReview(ctx *gin.Context, input *capstone_group_dto.FeedbackCapstoneGroupReviewInput) error {
	// TODO: Check current user is teacher belong to evaluation committee before update feedback
	var capstoneGroupReview model.CapstoneGroupReview

	if err := global.Db.Model(model.CapstoneGroupReview{}).
		Where("id = ? AND capstone_group_id = ?", input.CapstoneGroupReviewID, input.CapstoneGroupID).
		First(&capstoneGroupReview).
		Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupReviewNotFound,
		}))
	}

	if err := global.Db.Model(model.CapstoneGroupReview{}).
		Where("id = ?", input.CapstoneGroupReviewID).
		Select("Feedback").
		Updates(&model.CapstoneGroupReview{
			Feedback: input.Feedback,
		}).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		}))
	}

	return nil
}

func (c *capstoneGroupReviewService) GetCapstoneGroupReview(ctx *gin.Context, id int64, capstoneGroupID int64) (*capstone_group_dto.CapstoneGroupReviewOutput, error) {
	var capstoneGroupReview model.CapstoneGroupReview

	if err := global.Db.Model(model.CapstoneGroupReview{}).Where("id = ? AND capstone_group_id = ?", id, capstoneGroupID).First(&capstoneGroupReview).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupReviewNotFound,
		}))
	}

	output := capstone_group_dto.ToCapstoneGroupReviewOutput(&capstoneGroupReview)

	return output, nil
}
