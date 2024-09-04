package service

import (
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/major_dto"
	"github.com/api/internal/repository"
	"github.com/gin-gonic/gin"
)

type IMajorService interface {
	GetListMajor(ctx *gin.Context, input major_dto.InputGetListMajor) (interface{}, error)
}

type majorService struct {
	majorRepository repository.IMajorRepository
}

func NewMajorService(majorRepository repository.IMajorRepository) IMajorService {
	return &majorService{
		majorRepository: majorRepository,
	}
}

// NOTE: Consider to switch to concurrent call instead sequential
func (s *majorService) GetListMajor(ctx *gin.Context, input major_dto.InputGetListMajor) (interface{}, error) {
	total, err := s.majorRepository.CountAllMajor(ctx)

	if err != nil {
		return nil, err
	}
	items, err := s.majorRepository.GetListMajor(ctx, repository.GetListMajorParams {
		Limit: int32(input.Limit),
		Offset: int32(input.Offset),
	})

	if err != nil {
		return nil, err
	}

	itemsMajorOutput := make([]major_dto.OutputMajor, len(items))
	for i, item := range items {
		itemsMajorOutput[i] = major_dto.ToMajorOutput(item)
	}

	return major_dto.OutputGetListMajor {
		Meta: dto.MetaPagination {
			CurrentPage: int(input.Page),
			Total: int(total),
		},
		Items: itemsMajorOutput,
	}, nil
}