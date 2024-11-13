package capstone_group_controller

import (
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// @Summary UpdateReportFilesCapstoneGroupReview
// @Description Update report files capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.UpdateReportFilesCapstoneGroupReviewInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-reviews [patch]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateReportFilesCapstoneGroupReview(ctx *gin.Context) {
	var input capstone_group_dto.UpdateReportFilesCapstoneGroupReviewInput
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
	err = cgc.capstoneGroupReview.UpdateReportFilesCapstoneGroupReview(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateCapstoneGroupReviewSuccess,
	}))
}

// @Summary FeedbackCapstoneGroupReview
// @Description Feedback capstone group review
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.FeedbackCapstoneGroupReviewInput true "data"
// @Router /capstone-groups/{capstone_group_id}/capstone-group-reviews/feedback [patch]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) FeedbackCapstoneGroupReview(ctx *gin.Context) {
	var input capstone_group_dto.FeedbackCapstoneGroupReviewInput
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
	err = cgc.capstoneGroupReview.FeedbackCapstoneGroupReview(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateCapstoneGroupReviewSuccess,
	}))
}

// @Summary GetCapstoneGroupReview
// @Description Get capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Router /capstone-groups/{capstone_group_id}/capstone-group-reviews/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.GetCapstoneGroupReviewSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroupReview(ctx *gin.Context) {
	capstoneGroupIDStr := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	capstoneGroupReviewIDStr := ctx.Param("id")
	capstoneGroupReviewID, err := strconv.Atoi(capstoneGroupReviewIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	output, err := cgc.capstoneGroupReview.GetCapstoneGroupReview(ctx, int64(capstoneGroupReviewID), int64(capstoneGroupID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}
