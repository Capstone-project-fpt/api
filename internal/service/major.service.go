package service

import (
	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/major_dto"
	"github.com/gin-gonic/gin"
)

type IMajorService interface {
	GetListMajor(ctx *gin.Context, input major_dto.GetListMajorInput) (interface{}, error)
	GetMajor(ctx *gin.Context, id int) (interface{}, error)
}

type majorService struct{}

func NewMajorService() IMajorService {
	return &majorService{}
}

func (s *majorService) GetListMajor(ctx *gin.Context, input major_dto.GetListMajorInput) (interface{}, error) {
	var total int64
	if err := global.Db.Model(&model.Major{}).Count(&total).Error; err != nil {
		return nil, err
	}

	var items []model.Major
	if err := global.Db.Omit("CreatedAt", "UpdatedAt").Model(&model.Major{}).
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsMajorOutput := make([]major_dto.MajorOutput, len(items))
	for i, item := range items {
		itemsMajorOutput[i] = major_dto.ToMajorOutput(item)
	}

	return major_dto.GetListMajorOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsMajorOutput,
	}, nil
}

func (s *majorService) GetMajor(ctx *gin.Context, id int) (interface{}, error) {
	var major model.Major
	if err := global.Db.Where("id = ?", id).First(&major).Error; err != nil {
		return nil, err
	}

	majorOutput := major_dto.ToMajorOutput(major)

	return majorOutput, nil
}
