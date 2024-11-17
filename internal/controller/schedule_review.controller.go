package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/schedule_review_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ScheduleReviewController struct {
	scheduleReviewService service.IScheduleReviewService
}

func NewScheduleReviewController(scheduleReviewService service.IScheduleReviewService) *ScheduleReviewController {
	return &ScheduleReviewController{
		scheduleReviewService: scheduleReviewService,
	}
}

// @Summary CreateScheduleReview
// @Description Create Schedule Review
// @Tags Schedule Review
// @Produce json
// @Param data body schedule_review_dto.CreateScheduleReviewInput true "data"
// @Router /schedule-reviews [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (s *ScheduleReviewController) CreateScheduleReview(ctx *gin.Context) {
	var input schedule_review_dto.CreateScheduleReviewInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, errors.New(message))
		return
	}

	if err := global.Validator.Struct(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := s.scheduleReviewService.CreateScheduleReview(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateScheduleReviewSuccess,
	}))
}

// @Summary UpdateScheduleReview
// @Description Update Schedule Review
// @Tags Schedule Review
// @Produce json
// @Param data body schedule_review_dto.UpdateScheduleReviewInput true "data"
// @Router /schedule-reviews [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (s *ScheduleReviewController) UpdateScheduleReview(ctx *gin.Context) {
	var input schedule_review_dto.UpdateScheduleReviewInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, errors.New(message))
		return
	}

	if err := global.Validator.Struct(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := s.scheduleReviewService.UpdateScheduleReview(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateScheduleReviewSuccess,
	}))
}

// @Summary DeleteScheduleReview
// @Description Delete Schedule Review
// @Tags Schedule Review
// @Produce json
// @Router /schedule-reviews/{schedule_review_id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (s *ScheduleReviewController) DeleteScheduleReview(ctx *gin.Context) {
	scheduleReviewIDStr := ctx.Param("schedule_review_id")
	scheduleReviewID, err := strconv.Atoi(scheduleReviewIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = s.scheduleReviewService.DeleteScheduleReview(ctx, int64(scheduleReviewID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteScheduleReviewSuccess,
	}))
}

// @Summary GetScheduleReviewDetail
// @Description Get Schedule Review Detail
// @Tags Schedule Review
// @Produce json
// @Router /schedule-reviews/{schedule_review_id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} schedule_review_dto.ScheduleReviewDetailSwaggerOutput
// @Security ApiKeyAuth
func (s *ScheduleReviewController) GetScheduleReviewDetail(ctx *gin.Context) {
	scheduleReviewIDStr := ctx.Param("schedule_review_id")
	scheduleReviewID, err := strconv.Atoi(scheduleReviewIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	output, err := s.scheduleReviewService.GetScheduleReviewDetail(ctx, int64(scheduleReviewID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary GetListScheduleReview
// @Description Get List Schedule Review
// @Tags Schedule Review
// @Produce json
// @Param start_time query string true "Start Time"
// @Param end_time query string true "End Time"
// @Param order_by query string false "Order By"
// @Param evaluation_committee_id query int false "Evaluation Committee ID"
// @Param capstone_group_id query int false "Capstone Group ID"
// @Router /schedule-reviews [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} schedule_review_dto.ListScheduleReviewDetailSwaggerOutput
// @Security ApiKeyAuth
func (s *ScheduleReviewController) GetListScheduleReview(ctx *gin.Context) {
	var input schedule_review_dto.GetListScheduleReviewInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	output, err := s.scheduleReviewService.GetListScheduleReview(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}
