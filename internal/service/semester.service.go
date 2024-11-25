package service

import (
	"errors"
	"time"

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
	GetCurrentSemester(ctx *gin.Context) (*semester_dto.SemesterOutput, error)
	GetListSemester(ctx *gin.Context, input *semester_dto.GetListSemestersInput) (*semester_dto.ListSemestersOutput, error)
	GetListSemestersWithCountGroup(ctx *gin.Context, input *semester_dto.GetListSemestersInput) (*semester_dto.ListSemestersOutputCount, error)
}

type semesterService struct{}

func NewSemesterService() ISemesterService {
	return &semesterService{}
}

func (s *semesterService) CreateSemester(ctx *gin.Context, input *semester_dto.CreateSemesterInput) error {
	var overlapSemester model.Semester

	if input.StartTime.After(input.EndTime) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterStartTimeNeedToBeBeforeSemesterEndTime,
		})
		return errors.New(message)
	}

	if err := global.Db.Model(model.Semester{}).
		Where(
			"start_time <= ? AND end_time >= ?",
			input.EndTime.UTC(),
			input.StartTime.UTC(),
		).First(&overlapSemester).Error; err == nil {
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

	if input.StartTime.After(input.EndTime) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterStartTimeNeedToBeBeforeSemesterEndTime,
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

func (s *semesterService) GetCurrentSemester(ctx *gin.Context) (*semester_dto.SemesterOutput, error) {
	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("start_time <= ? AND end_time >= ?", time.Now(), time.Now()).First(&semester).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CurrentSemesterNotFound,
		})
		return nil, errors.New(message)
	}

	semesterOutput := semester_dto.ToSemesterOutput(&semester)

	return &semesterOutput, nil
}

func (s *semesterService) GetListSemester(ctx *gin.Context, input *semester_dto.GetListSemestersInput) (*semester_dto.ListSemestersOutput, error) {
	var total int64
	var items []model.Semester
	query := global.Db.Model(model.Semester{})

	if err := global.Db.Model(model.Semester{}).Count(&total).Error; err != nil {
		return nil, err
	}

	if input.OrderBy != "" {
		query.Order("semesters.start_time " + input.OrderBy)
	}

	if err := query.
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
func (s *semesterService) GetListSemestersWithCountGroup(ctx *gin.Context, input *semester_dto.GetListSemestersInput) (*semester_dto.ListSemestersOutputCount, error) {
    var items []semester_dto.SemesterOutputCount
    var total int64

    err := global.Db.Model(&model.Semester{}).
        Select(`
            "semesters"."id",
            "semesters"."name",
            "semesters"."start_time",
            "semesters"."end_time",
            COUNT(DISTINCT "capstone_groups"."id") AS "capstone_groups",
            COUNT(DISTINCT "evaluation_committees"."id") AS "evaluation_committees"
        `).
        Joins(`LEFT JOIN "capstone_groups" ON "capstone_groups"."semester_id" = "semesters"."id"`).
        Joins(`LEFT JOIN "evaluation_committees" ON "evaluation_committees"."semester_id" = "semesters"."id"`).
        Group(`"semesters"."id"`).
        Order(`"semesters"."start_time" DESC`).
		Limit (int(input.Limit)).
		Offset(int(input.Offset)).
        Find(&items).Error

    if err != nil {
        return nil, err
    }

    err = global.Db.Model(&model.Semester{}).Count(&total).Error
    if err != nil {
        return nil, err
    }

    return &semester_dto.ListSemestersOutputCount{
        Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
            Total:       int(total),
        },
        Items: items,
    }, nil
}
