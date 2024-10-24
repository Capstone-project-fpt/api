package capstone_group_controller

import (
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/capstone_group_topic_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// @Summary CreateCapstoneGroupTopic
// @Description Create capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_topic_dto.CreateCapstoneGroupTopicInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) CreateCapstoneGroupTopic(ctx *gin.Context) {
	var input capstone_group_topic_dto.CreateCapstoneGroupTopicInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.CapstoneGroupID = int64(capstoneGroupID)
	err = cgc.capstoneGroupTopicService.CreateCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateCapstoneGroupTopicSuccess,
	}))
}

// @Summary UpdateCapstoneGroupTopic
// @Description Update capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_topic_dto.CreateCapstoneGroupTopicInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id} [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateCapstoneGroupTopic(ctx *gin.Context) {
	var input capstone_group_topic_dto.UpdateCapstoneGroupTopicInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.CapstoneGroupTopicID = int64(id)
	input.CapstoneGroupID = int64(capstoneGroupID)
	err = cgc.capstoneGroupTopicService.UpdateCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateCapstoneGroupTopicSuccess,
	}))
}

// @Summary DeleteCapstoneGroupTopic
// @Description Delete capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) DeleteCapstoneGroupTopic(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input := service.DeleteCapstoneGroupTopicInput{
		CapstoneGroupTopicID: int64(id),
		CapstoneGroupID:      int64(capstoneGroupID),
	}
	err = cgc.capstoneGroupTopicService.DeleteCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteCapstoneGroupTopicSuccess,
	}))
}

// @Summary ReviewCapstoneGroupTopic
// @Description Review capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_topic_dto.ReviewCapstoneGroupTopicInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id}/teacher-reviews [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) ReviewCapstoneGroupTopic(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input capstone_group_topic_dto.ReviewCapstoneGroupTopicInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	if err := global.Validator.Struct(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
	}
	input.CapstoneGroupTopicID = int64(id)
	input.CapstoneGroupID = int64(capstoneGroupID)

	err = cgc.capstoneGroupTopicService.ReviewCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.ReviewCapstoneGroupTopicSuccess,
	}))
}

// @Summary FeedbackCapstoneGroupTopic
// @Description Feedback capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_topic_dto.FeedbackCapstoneGroupTopicInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id}/feedbacks [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) FeedbackCapstoneGroupTopic(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input capstone_group_topic_dto.FeedbackCapstoneGroupTopicInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}
	if err := global.Validator.Struct(input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	input.CapstoneGroupTopicID = int64(id)
	input.CapstoneGroupID = int64(capstoneGroupID)

	err = cgc.capstoneGroupTopicService.FeedbackCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.FeedbackCapstoneGroupTopicSuccess,
	}))
}

// @Summary UpdateFeedbackCapstoneGroupTopic
// @Description Update feedback capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_topic_dto.UpdateFeedbackCapstoneGroupTopicInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id}/feedbacks/{feedback_id} [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateFeedbackCapstoneGroupTopic(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	capstoneGroupTopicIDStr := ctx.Param("id")
	capstoneGroupTopicID, err := strconv.Atoi(capstoneGroupTopicIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	feedbackIDStr := ctx.Param("feedback_id")
	feedbackID, err := strconv.Atoi(feedbackIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input capstone_group_topic_dto.UpdateFeedbackCapstoneGroupTopicInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	if err := global.Validator.Struct(input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}
	input.FeedbackID = int64(feedbackID)
	input.CapstoneGroupTopicID = int64(capstoneGroupTopicID)
	input.CapstoneGroupID = int64(capstoneGroupID)

	err = cgc.capstoneGroupTopicService.UpdateFeedbackCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateFeedbackCapstoneGroupTopicSuccess,
	}))
}

// @Summary DeleteFeedbackCapstoneGroupTopic
// @Description Delete feedback capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id}/feedbacks/{feedback_id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) DeleteFeedbackCapstoneGroupTopic(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	capstoneGroupTopicIDStr := ctx.Param("id")
	capstoneGroupTopicID, err := strconv.Atoi(capstoneGroupTopicIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	feedbackIDStr := ctx.Param("feedback_id")
	feedbackID, err := strconv.Atoi(feedbackIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input service.DeleteFeedbackCapstoneGroupTopicInput
	input.FeedbackID = int64(feedbackID)
	input.CapstoneGroupTopicID = int64(capstoneGroupTopicID)
	input.CapstoneGroupID = int64(capstoneGroupID)

	err = cgc.capstoneGroupTopicService.DeleteFeedbackCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteFeedbackCapstoneGroupTopicSuccess,
	}))
}

// @Summary GetCapstoneGroupTopic
// @Description Get capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_topic_dto.GetCapstoneTopicGroupSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroupTopic(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input := service.GetCapstoneGroupTopicInput{
		CapstoneGroupTopicID: int64(id),
		CapstoneGroupID:      int64(capstoneGroupID),
	}
	capstoneGroupTopic, err := cgc.capstoneGroupTopicService.GetCapstoneGroupTopic(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, capstoneGroupTopic)
}

// @Summary GetListCapstoneGroupTopic
// @Description Get list capstone group topic
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param order_by query string false "Order by"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_topic_dto.ListCapstoneGroupTopicsOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetListCapstoneGroupTopic(ctx *gin.Context) {
	var input capstone_group_topic_dto.GetListCapstoneGroupTopicInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.CapstoneGroupID = int64(capstoneGroupID)
	capstoneGroupTopic, err := cgc.capstoneGroupTopicService.GetListCapstoneGroupTopics(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, capstoneGroupTopic)
}

// @Summary GetCapstoneGroupTopicFeedback
// @Description Get capstone group topic feedback
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id}/feedbacks/{feedback_id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_topic_dto.GetCapstoneGroupTopicFeedbackSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroupTopicFeedback(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	_, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	feedbackIDStr := ctx.Param("feedback_id")
	feedbackID, err := strconv.Atoi(feedbackIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input := service.GetCapstoneGroupTopicFeedbackInput{
		CapstoneGroupTopicID: int64(id),
		FeedbackID:           int64(feedbackID),
	}
	capstoneGroupTopicFeedback, err := cgc.capstoneGroupTopicService.GetCapstoneGroupTopicFeedback(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, capstoneGroupTopicFeedback)
}

// @Summary GetListCapstoneGroupTopicFeedback
// @Description Get list capstone group topic feedback
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param order_by query string false "Order by"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-topics/{id}/feedbacks [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_topic_dto.ListCapstoneGroupTopicFeedbackOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetListCapstoneGroupTopicFeedback(ctx *gin.Context) {
	var input capstone_group_topic_dto.GetListCapstoneGroupTopicFeedbackInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	_, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	input.CapstoneGroupTopicID = int64(id)
	capstoneGroupTopicFeedbacks, err := cgc.capstoneGroupTopicService.GetListCapstoneGroupTopicFeedbacks(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, capstoneGroupTopicFeedbacks)
}