package capstone_group_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_dto"
	capstone_group_service "github.com/api/internal/service/capstone_group"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type CapstoneGroupController struct {
	capstoneGroupService      capstone_group_service.ICapstoneGroupService
	capstoneGroupTopicService capstone_group_service.ICapstoneGroupTopicService
	capstoneGroupReview       capstone_group_service.ICapstoneGroupReviewService
}

func NewCapstoneGroupController(
	capstoneGroupService capstone_group_service.ICapstoneGroupService,
	capstoneGroupTopicService capstone_group_service.ICapstoneGroupTopicService,
	capstoneGroupReview capstone_group_service.ICapstoneGroupReviewService,
) *CapstoneGroupController {
	return &CapstoneGroupController{
		capstoneGroupService:      capstoneGroupService,
		capstoneGroupTopicService: capstoneGroupTopicService,
		capstoneGroupReview:       capstoneGroupReview,
	}
}

// @Summary CreateCapstoneGroup
// @Description Create capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.CreateCapstoneGroupInput true "data"
// @Router /capstone-groups [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.GetCapstoneGroupSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) CreateCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.CreateCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	capstoneGroupOutput, err := cgc.capstoneGroupService.CreateCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, capstoneGroupOutput)
}

// @Summary UpdateCapstoneGroup
// @Description Update capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.UpdateCapstoneGroupInput true "data"
// @Router /capstone-groups [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.UpdateCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err := cgc.capstoneGroupService.UpdateCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateCapstoneGroupSuccess,
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
}

// @Summary InviteMentorToCapstoneGroup
// @Description Invite mentor to capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.InviteMentorToCapstoneGroupInput true "data"
// @Router /capstone-groups/{capstone_group_id}/mentors [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) InviteMentorToCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.InviteMentorToCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	idStr := ctx.Param("capstone_group_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.CapstoneGroupID = int64(id)

	err = cgc.capstoneGroupService.InviteMentorToCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.SendInviteToMentorSuccess,
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
}

// @Summary ResponseInviteMentorToCapstoneGroup
// @Description Response invite mentor to capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.ResponseInviteMentorToCapstoneGroupInput true "data"
// @Router /capstone-groups/{capstone_group_id}/mentors/invitations [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) ResponseInviteMentorToCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.ResponseInviteMentorToCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	if err := global.Validator.Struct(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	idStr := ctx.Param("capstone_group_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.CapstoneGroupID = int64(id)

	err = cgc.capstoneGroupService.ResponseInviteMentorToCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.ResponseInviteMentorToCapstoneGroupSuccess,
		TemplateData: map[string]interface{}{
			"Status": input.Status,
		},
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
}

// @Summary GetListCapstoneGroups
// @Description Get list capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param semester_id query int false "SemesterID"
// @Router /capstone-groups [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ListCapstoneGroupOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetListCapstoneGroups(ctx *gin.Context) {
	var input capstone_group_dto.GetListCapstoneGroupInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	result, err := cgc.capstoneGroupService.GetListCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, result)
}

// @Summary GetCapstoneGroup
// @Description Get capstone group
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Router /capstone-groups/{capstone_group_id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.GetCapstoneGroupSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroup(ctx *gin.Context) {
	idParam := ctx.Param("capstone_group_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	output, err := cgc.capstoneGroupService.GetCapstoneGroup(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary GetMentorAndListMemberCapstoneGroup
// @Description Get mentor and list member capstone group
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Router /capstone-groups/{capstone_group_id}/members [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.MentorAndListMemberCapstoneGroupSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetMentorAndListMemberCapstoneGroup(ctx *gin.Context) {
	idParam := ctx.Param("capstone_group_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	output, err := cgc.capstoneGroupService.GetMentorAndListMemberCapstoneGroup(ctx, int64(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary GetListInvitationMentorCapstoneGroups
// @Description Get list invitation mentor capstone group
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Router /capstone-groups/{capstone_group_id}/mentors/invitations [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ListInvitationMentorCapstoneGroupOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetListInvitationMentorCapstoneGroups(ctx *gin.Context) {
	idParam := ctx.Param("capstone_group_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	var input capstone_group_dto.GetListInviteMentorToCapstoneGroupInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	input.CapstoneGroupID = int64(id)

	output, err := cgc.capstoneGroupService.GetListInvitationMentorCapstoneGroups(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary GetListStudentHaveCapstoneGroup
// @Description Get list student have capstone group
// @Tags Capstone Group
// @Produce json
// @Param semester_id path int true "semester_id"
// @Router /capstone-groups/semesters/{semester_id}/students [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ListStudentHaveCapstoneGroupSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetListStudentHaveCapstoneGroup(ctx *gin.Context) {
	idParam := ctx.Param("semester_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	output, err := cgc.capstoneGroupService.GetListStudentHaveCapstoneGroup(ctx, int64(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary UpdateCapstoneGroupStudent
// @Description Update Capstone Group Member
// @Tags Capstone Group
// @Produce json
// @Param id path int true "capstone_group_id"
// @Param data body capstone_group_dto.UpdateCapstoneGroupStudentInput true "data"
// @Router /capstone-groups/{capstone_group_id}/members [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.UpdateCapstoneGroupStudentInput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateCapstoneGroupStudent(ctx *gin.Context) {
	idParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.UpdateCapstoneGroupStudentInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})
		response.ErrorResponse(ctx, http.StatusBadRequest, errors.New(message))
		return
	}

	input.ID = int64(capstoneGroupID)

	if err := global.Validator.Struct(input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := cgc.capstoneGroupService.UpdateCapstoneGroupStudent(ctx, &input); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateCapstoneGroupStudentSuccess,
	}))
}

// @Summary CreateCapstoneGroupReportDocument
// @Description CreateCapstoneGroupReportDocument
// @Tags Capstone Group
// @Produce json
// @Param id path int true "capstone_group_id"
// @Param data body capstone_group_dto.CreateCapstoneGroupReportDocumentInput true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) CreateCapstoneGroupReportDocument(ctx *gin.Context) {
	idParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.CreateCapstoneGroupReportDocumentInput
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

	input.CapstoneGroupID = int64(capstoneGroupID)

	if err := cgc.capstoneGroupService.CreateCapstoneGroupReportDocument(ctx, &input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateCapstoneGroupReportDocumentSuccess,
	}))
}

// @Summary UpdateCapstoneGroupReportDocument
// @Description UpdateCapstoneGroupReportDocument
// @Tags Capstone Group
// @Produce json
// @Param id path int true "capstone_group_id"
// @Param data body capstone_group_dto.UpdateCapstoneGroupReportDocumentInput true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) UpdateCapstoneGroupReportDocument(ctx *gin.Context) {
	idParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.UpdateCapstoneGroupReportDocumentInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})
		response.ErrorResponse(ctx, http.StatusBadRequest, errors.New(message))
		return
	}

	input.CapstoneGroupID = int64(capstoneGroupID)

	if err := cgc.capstoneGroupService.UpdateCapstoneGroupReportDocument(ctx, &input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateCapstoneGroupReportDocumentSuccess,
	}))
}

// @Summary DeleteCapstoneGroupReportDocument
// @Description DeleteCapstoneGroupReportDocument
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Router /capstone-groups/{capstone_group_id}/report-documents/{report_document_id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) DeleteCapstoneGroupReportDocument(ctx *gin.Context) {
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

	if err := cgc.capstoneGroupService.DeleteCapstoneGroupReportDocument(ctx, int64(reportDocumentID), int64(capstoneGroupID)); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteCapstoneGroupReportDocumentSuccess,
	}))
}

// @Summary GetCapstoneGroupReportDocument
// @Description GetCapstoneGroupReportDocument
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Router /capstone-groups/{capstone_group_id}/report-documents/{report_document_id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ReportDocumentSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroupReportDocument(ctx *gin.Context) {
	reportDocumentIDParam := ctx.Param("report_document_id")
	reportDocumentID, err := strconv.Atoi(reportDocumentIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	output, err := cgc.capstoneGroupService.GetCapstoneGroupReportDocument(ctx, int64(reportDocumentID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary GetCapstoneGroupReportDocument
// @Description GetCapstoneGroupReportDocument
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Router /capstone-groups/{capstone_group_id}/report-documents [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ReportDocumentsSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroupReportDocuments(ctx *gin.Context) {
	capstoneGroupIDParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	output, err := cgc.capstoneGroupService.GetCapstoneGroupReportDocuments(ctx, int64(capstoneGroupID))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary MentorUpdateStudentScoreForReportDocument
// @Description MentorUpdateStudentScoreForReportDocument
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param data body capstone_group_dto.MentorUpdateStudentScoreForReportDocument true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents/scores [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) MentorUpdateStudentScoreForReportDocument(ctx *gin.Context) {
	capstoneGroupIDParam := ctx.Param("capstone_group_id")
	capstoneGroupID, err := strconv.Atoi(capstoneGroupIDParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	var input capstone_group_dto.MentorUpdateStudentScoreForReportDocument
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	input.CapstoneGroupID = int64(capstoneGroupID)

	if err := cgc.capstoneGroupService.MentorUpdateStudentScoreForReportDocument(ctx, &input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.MentorUpdateStudentScoreForReportDocumentSuccess,
	}))
}

// @Summary GetReportDocumentStudentsScore
// @Description GetReportDocumentStudentsScore
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param report_document_id path int true "report_document_id"
// @Router /capstone-groups/{capstone_group_id}/report-documents/scores/{report_document_id} [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.ReportDocumentStudentsScoreSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetReportDocumentStudentsScore(ctx *gin.Context) {
	idParam := ctx.Param("report_document_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	output, err := cgc.capstoneGroupService.GetReportDocumentStudentsScore(ctx, int64(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary AdminUpdateStudentScoreForReportDocument
// @Description AdminUpdateStudentScoreForReportDocument
// @Tags Capstone Group
// @Produce json
// @Param capstone_group_id path int true "capstone_group_id"
// @Param data body capstone_group_dto.AdminUpdateStudentScoreForReportDocument true "data"
// @Router /capstone-groups/{capstone_group_id}/report-documents/scores [patch]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) AdminUpdateStudentScoreForReportDocument(ctx *gin.Context) {
	var input capstone_group_dto.AdminUpdateStudentScoreForReportDocument
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		}))
		return
	}

	err := cgc.capstoneGroupService.AdminUpdateStudentScoreForReportDocument(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.AdminUpdateStudentScoreForReportDocumentSuccess,
	}))
}
