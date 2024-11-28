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

// @Summary CommentReport
// @Description Student or Mentor of capstone group comment on report
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Param data body capstone_group_dto.CommentReportInput true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents/{report_document_id}/comments [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) CommentReport(ctx *gin.Context) {
	capstoneGroupIDParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	reportDocumentIDParam := ctx.Param("report_document_id")
	reportDocumentID, err := strconv.Atoi(reportDocumentIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.CommentReportInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	input.CapstoneGroupID = int64(capstoneGroupID)
	input.ReportDocumentID = int64(reportDocumentID)

	err = cgc.capstoneGroupService.CommentReport(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CommentSuccess,
	}))
}

// @Summary UpdateCommentReport
// @Description Student or Mentor of capstone group update comment on report
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Param data body capstone_group_dto.UpdateCommentReportInput true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents/{report_document_id}/comments [patch]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateCommentReport(ctx *gin.Context) {
	capstoneGroupIDParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	reportDocumentIDParam := ctx.Param("report_document_id")
	reportDocumentID, err := strconv.Atoi(reportDocumentIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.UpdateCommentReportInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	input.CapstoneGroupID = int64(capstoneGroupID)
	input.ReportDocumentID = int64(reportDocumentID)

	err = cgc.capstoneGroupService.UpdateCommentReport(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateCommentSuccess,
	}))
}

// @Summary DeleteCommentReport
// @Description Student or Mentor of capstone group delete comment on report
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Param data body capstone_group_dto.DeleteCommentReportInput true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents/{report_document_id}/comments [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) DeleteCommentReport(ctx *gin.Context) {
	capstoneGroupIDParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	reportDocumentIDParam := ctx.Param("report_document_id")
	reportDocumentID, err := strconv.Atoi(reportDocumentIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.DeleteCommentReportInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	input.CapstoneGroupID = int64(capstoneGroupID)
	input.ReportDocumentID = int64(reportDocumentID)

	err = cgc.capstoneGroupService.DeleteCommentReport(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteCommentSuccess,
	}))
}

// @Summary GetListReportComments
// @Description Get list report comments
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Param order query string false "order"
// @Router /capstone-groups/{capstone_group_id}/report-documents/{report_document_id}/comments [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ReportCommentsSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetListReportComments(ctx *gin.Context) {
	capstoneGroupIDParam := ctx.Param("capstone_group_id")
	_, err := strconv.Atoi(capstoneGroupIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	reportDocumentIDParam := ctx.Param("report_document_id")
	reportDocumentID, err := strconv.Atoi(reportDocumentIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.GetListCommentReportInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	input.ReportDocumentID = int64(reportDocumentID)

	output, err := cgc.capstoneGroupService.GetListReportComments(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}
