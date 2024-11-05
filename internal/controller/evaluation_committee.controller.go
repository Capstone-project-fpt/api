package controller

import (
	"net/http"
	"strconv"

	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
)

type EvaluationCommitteeController struct {
	evaluationCommitteeService service.IEvaluationCommitteeService
}

func NewEvaluationCommitteeController(evaluationCommitteeService service.IEvaluationCommitteeService) *EvaluationCommitteeController {
	return &EvaluationCommitteeController{
		evaluationCommitteeService,
	}
}

// @Summary GetEvaluationCommittee
// @Description Get Evaluation Committee
// @Tags Evaluation Committee
// @Produce json
// @Param id path int true "id"
// @Router /evaluation-committees/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} evaluation_committee_dto.EvaluationCommitteeSwaggerOutput
// @Security ApiKeyAuth
func (ecc *EvaluationCommitteeController) GetEvaluationCommittee(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	output, err := ecc.evaluationCommitteeService.GetEvaluationCommittee(ctx, int64(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}
