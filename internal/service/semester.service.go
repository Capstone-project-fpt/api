package service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/semester_dto"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ISemesterService interface {
	CreateSemester(ctx *gin.Context, input *semester_dto.CreateSemesterInput) error
	UpdateSemester(ctx *gin.Context, input *semester_dto.UpdateSemesterInput) error
	DeleteSemester(ctx *gin.Context, id int) error
	GetSemester(ctx *gin.Context, id int) (*semester_dto.SemesterOutput, error)
	GetListSemester(ctx *gin.Context, input *semester_dto.GetListSemestersInput) (*semester_dto.ListSemestersOutput, error)
}

type semesterService struct{}

func NewSemesterService() ISemesterService {
	return &semesterService{}
}

func (s *semesterService) CreateSemester(ctx *gin.Context, input *semester_dto.CreateSemesterInput) error {
	var overlapSemester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("start_time < ? AND end_time > ?", input.EndTime, input.StartTime).First(&overlapSemester).Error; err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterOverlap,
		})
		return errors.New(message)
	}

	if err := global.Db.Model(model.Semester{}).Create(&model.Semester{
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Name:      input.Name,
	}).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return errors.New(message)
	}

	return nil
}

func (s *semesterService) UpdateSemester(ctx *gin.Context, input *semester_dto.UpdateSemesterInput) error {
	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", input.ID).First(&semester).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		})
		return errors.New(message)
	}

	var overlapSemester model.Semester
	if err := global.Db.Model(model.Semester{}).
		Where("start_time < ? AND end_time > ? AND id != ?", input.EndTime, input.StartTime, input.ID).
		First(&overlapSemester).Error; err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterOverlap,
		})
		return errors.New(message)
	}

	if err := global.Db.Model(model.Semester{}).Where("id = ?", input.ID).Updates(&model.Semester{
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Name:      input.Name,
	}).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return errors.New(message)
	}

	return nil
}

func (s *semesterService) DeleteSemester(ctx *gin.Context, id int) error {
	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", id).First(&semester).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		})
		return errors.New(message)
	}

	if err := global.Db.Model(model.Semester{}).Where("id = ?", id).Delete(&model.Semester{}).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return errors.New(message)
	}

	return nil
}

func (s *semesterService) GetSemester(ctx *gin.Context, id int) (*semester_dto.SemesterOutput, error) {
	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", id).First(&semester).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		})
		return nil, errors.New(message)
	}

	semesterOutput := semester_dto.ToSemesterOutput(&semester)

	return &semesterOutput, nil
}

func (s *semesterService) GetListSemester(ctx *gin.Context, input *semester_dto.GetListSemestersInput) (*semester_dto.ListSemestersOutput, error) {
	var total int64
	var items []model.Semester

	if err := global.Db.Model(model.Semester{}).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := global.Db.Model(model.Semester{}).
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsSemesterOutput := make([]semester_dto.SemesterOutput, len(items))
	for i, item := range items {
		itemsSemesterOutput[i] = semester_dto.ToSemesterOutput(&item)
	}

	return &semester_dto.ListSemestersOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsSemesterOutput,
	}, nil
}
