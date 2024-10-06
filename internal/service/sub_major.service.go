package service

import (
	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/sub_major_dto"
	"github.com/gin-gonic/gin"
)

type ISubMajorService interface {
	GetListSubMajor(ctx *gin.Context, input sub_major_dto.GetListSubMajorInput) (interface{}, error)
	GetSubMajor(ctx *gin.Context, id int) (interface{}, error)
}

type subMajorService struct{}

func NewSubMajorService() ISubMajorService {
	return &subMajorService{}
}

func (s *subMajorService) GetListSubMajor(ctx *gin.Context, input sub_major_dto.GetListSubMajorInput) (interface{}, error) {
	var total int64
	var items []model.SubMajor
	getTotalQuery := global.Db.Model(&model.SubMajor{})
	getSubMajorsQuery := global.Db.Model(&model.SubMajor{})

	if input.MajorID != 0 {
		getTotalQuery = getTotalQuery.Where("major_id = ?", input.MajorID)
		getSubMajorsQuery = getSubMajorsQuery.Where("major_id = ?", input.MajorID)
	}

	if err := getTotalQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := getSubMajorsQuery.Omit("CreatedAt", "UpdatedAt").
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsSubMajorOutput := make([]sub_major_dto.SubMajorOutput, len(items))
	for i, item := range items {
		itemsSubMajorOutput[i] = sub_major_dto.ToSubMajorOutput(item)
	}

	return sub_major_dto.GetListSubMajorOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsSubMajorOutput,
	}, nil
}

func (s *subMajorService) GetSubMajor(ctx *gin.Context, id int) (interface{}, error) {
	var subMajor model.SubMajor
	if err := global.Db.Model(&model.SubMajor{}).Where("id = ?", id).First(&subMajor).Error; err != nil {
		return nil, err
	}

	subMajorOutput := sub_major_dto.ToSubMajorOutput(subMajor)

	return subMajorOutput, nil
}
