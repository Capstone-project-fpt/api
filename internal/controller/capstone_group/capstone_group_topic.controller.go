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
