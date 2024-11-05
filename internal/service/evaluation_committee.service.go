package service

import (
	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/dto/evaluation_committee_dto"
	"github.com/gin-gonic/gin"
)

type IEvaluationCommitteeService interface {
	GetEvaluationCommittee(ctx *gin.Context, id int64) (*evaluation_committee_dto.EvaluationCommitteeOutput, error)
}

type evaluationCommitteeService struct{}

func NewEvaluationCommitteeService() IEvaluationCommitteeService {
	return &evaluationCommitteeService{}
}

func (s *evaluationCommitteeService) GetEvaluationCommittee(ctx *gin.Context, id int64) (*evaluation_committee_dto.EvaluationCommitteeOutput, error) {
	var evaluationCommittee model.EvaluationCommittee

	query := global.Db.Model(&model.EvaluationCommittee{}).Where("id = ?", id).First(&evaluationCommittee)

	if err := query.Error; err != nil {
		return nil, err
	}

	output := evaluation_committee_dto.ToEvaluationCommitteeOutput(&evaluationCommittee)

	return output, nil
}
