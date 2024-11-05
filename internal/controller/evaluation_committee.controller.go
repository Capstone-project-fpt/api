package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/evaluation_committee_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type EvaluationCommitteeController struct {
	evaluationCommitteeService service.IEvaluationCommitteeService
}

func NewEvaluationCommitteeController(evaluationCommitteeService service.IEvaluationCommitteeService) *EvaluationCommitteeController {
	return &EvaluationCommitteeController{
		evaluationCommitteeService,
	}
}

// @Summary CreateEvaluationCommittee
// @Description Create Evaluation Committee
// @Tags Evaluation Committee
// @Produce json
// @Param data body evaluation_committee_dto.CreateEvaluationCommitteeInput true "data"
// @Router /evaluation-committees [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ecc *EvaluationCommitteeController) CreateEvaluationCommittee(ctx *gin.Context) {
	var input evaluation_committee_dto.CreateEvaluationCommitteeInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, errors.New(message))
		return
	}

	if err := global.Validator.Struct(input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := ecc.evaluationCommitteeService.CreateEvaluationCommittee(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateEvaluationCommitteeSuccess,
	}))
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
