package controller

import (
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type CapstoneGroupController struct {
	capstoneGroupService service.ICapstoneGroupService
}

func NewCapstoneGroupController(capstoneGroupService service.ICapstoneGroupService) *CapstoneGroupController {
	return &CapstoneGroupController{
		capstoneGroupService: capstoneGroupService,
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
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) CreateCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.CreateCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err := cgc.capstoneGroupService.CreateCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateCapstoneGroupSuccess,
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
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
// @Router /capstone-groups/{id}/mentors [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) InviteMentorToCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.InviteMentorToCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	idStr := ctx.Param("id")
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

// @Summary AcceptInviteMentorToCapstoneGroup
// @Description Accept invite mentor to capstone group
// @Tags Capstone Group
// @Accept json
// @Produce json
// @Param data body capstone_group_dto.AcceptInviteMentorToCapstoneGroupInput true "data"
// @Router /capstone-groups/{id}/mentors/invitation [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) AcceptInviteMentorToCapstoneGroup(ctx *gin.Context) {
	var input capstone_group_dto.AcceptInviteMentorToCapstoneGroupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.CapstoneGroupID = int64(id)

	err = cgc.capstoneGroupService.AcceptInviteMentorToCapstoneGroup(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.AcceptInviteMentorToCapstoneGroupSuccess,
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
// @Param id path int true "id"
// @Router /capstone-groups/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} capstone_group_dto.GetCapstoneGroupSwaggerOutput
// @Security ApiKeyAuth
func (cgc *CapstoneGroupController) GetCapstoneGroup(ctx *gin.Context) {
	idParam := ctx.Param("id")
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
