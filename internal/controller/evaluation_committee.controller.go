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
	util "github.com/api/pkg/utils"
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

// @Summary UpdateEvaluationCommittee
// @Description Update Evaluation Committee
// @Tags Evaluation Committee
// @Produce json
// @Param data body evaluation_committee_dto.UpdateEvaluationCommitteeInput true "data"
// @Router /evaluation-committees [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ecc *EvaluationCommitteeController) UpdateEvaluationCommittee(ctx *gin.Context) {
	var input evaluation_committee_dto.UpdateEvaluationCommitteeInput
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

	err := ecc.evaluationCommitteeService.UpdateEvaluationCommittee(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateEvaluationCommitteeSuccess,
	}))
}

// @Summary DeleteEvaluationCommittee
// @Description Delete Evaluation Committee
// @Tags Evaluation Committee
// @Produce json
// @Router /evaluation-committees/{id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ecc *EvaluationCommitteeController) DeleteEvaluationCommittee(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ecc.evaluationCommitteeService.DeleteEvaluationCommittee(ctx, int64(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteEvaluationCommitteeSuccess,
	}))
}

// @Summary GetEvaluationCommittee
// @Description Get Evaluation Committee
// @Tags Evaluation Committee
// @Produce json
// @Param id path int true "id"
// @Router /evaluation-committees/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} evaluation_committee_dto.EvaluationCommitteeWithFullInfoSwaggerOutput
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

// @Summary GetListEvaluationCommittee
// @Description Get list evaluation committee
// @Tags Evaluation Committee
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param order_by query string false "Order by"
// @Param semester_id query int false "SemesterID"
// @Router /evaluation-committees [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} major_dto.GetListMajorOutput
func (ecc *EvaluationCommitteeController) GetListEvaluationCommittee(ctx *gin.Context) {
	var input evaluation_committee_dto.GetListEvaluationCommitteeInput
	localizer := global.Localizer
	if err := ctx.ShouldBindQuery(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	output, err := ecc.evaluationCommitteeService.GetListEvaluationCommittee(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary GetListTeachersHaveEvaluationCommitteeGroup
// @Description Get list teachers have evaluation committee group
// @Tags Evaluation Committee
// @Accept json
// @Produce json
// @Router /evaluation-committees/semesters/{semester_id}/teachers [get]
// @Param semester_id path int true "semester_id"
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} evaluation_committee_dto.ListTeachersHaveEvaluationCommitteeGroupSwaggerOutput
// @Security ApiKeyAuth
func (ecc *EvaluationCommitteeController) GetListTeachersHaveEvaluationCommitteeGroup(ctx *gin.Context) {
	idParam := ctx.Param("semester_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	output, err := ecc.evaluationCommitteeService.GetListTeachersHaveEvaluationCommitteeGroup(ctx, int64(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}
