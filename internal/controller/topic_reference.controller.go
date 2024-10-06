package controller

import (
	"net/http"
	"strconv"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/topic_reference_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TopicReferenceController struct {
	topicReferenceService service.ITopicReferenceService
}

func NewTopicReferenceController(topicReferenceService service.ITopicReferenceService) *TopicReferenceController {
	return &TopicReferenceController{
		topicReferenceService: topicReferenceService,
	}
}

// @Summary GetTopicReference
// @Description Get topic reference
// @Tags topic reference
// @Produce json
// @Param id path int true "id"
// @Router /topic-references/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} topic_reference_dto.GetTopicReferenceSwaggerOutput
// @Security ApiKeyAuth
func (trc *TopicReferenceController) GetTopicReference(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	topicReference, err := trc.topicReferenceService.GetTopicReference(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, topicReference)
}

// @Summary GetListTopicReferences
// @Description Get list of topic references
// @Tags topic reference
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param teacher_ids query []int false "TeacherIDs" collectionFormat(multi)
// @Router /topic-references [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} topic_reference_dto.ListTopicReferenceOutput
// @Security ApiKeyAuth
func (trc *TopicReferenceController) GetListTopicReferences(ctx *gin.Context) {
	var input topic_reference_dto.GetListTopicReferencesInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	result, err := trc.topicReferenceService.GetListTopicReferences(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, result)
}

// @Summary TeacherCreateTopicReference
// @Description Teacher create topic reference
// @Tags topic reference
// @Accept json
// @Produce json
// @Param data body topic_reference_dto.TeacherCreateTopicReferenceInput true "data"
// @Router /topic-references/teachers [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (trc *TopicReferenceController) TeacherCreateTopicReference(ctx *gin.Context) {
	var input topic_reference_dto.TeacherCreateTopicReferenceInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	userContext := context_util.GetUserContext(ctx)
	if userContext == nil {
		response.ErrorResponse(ctx, http.StatusNotFound, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
		return
	}

	err := trc.topicReferenceService.TeacherCreateTopicReference(
		ctx,
		service.CreateTopicReferenceInput{
			Name: input.Name,
			Path: input.Path,
		},
		userContext,
	)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, nil)
}

// @Summary TeacherUpdateTopicReference
// @Description Teacher update topic reference
// @Tags topic reference
// @Accept json
// @Produce json
// @Param data body topic_reference_dto.TeacherUpdateTopicReferenceInput true "data"
// @Router /topic-references/teachers [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (trc *TopicReferenceController) TeacherUpdateTopicReference(ctx *gin.Context) {
	var input topic_reference_dto.TeacherUpdateTopicReferenceInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	userContext := context_util.GetUserContext(ctx)
	if userContext == nil {
		response.ErrorResponse(ctx, http.StatusNotFound, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
		return
	}

	var teacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Select("id").Where("user_id = ?", userContext.ID).First(&teacher).Error; err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
		return
	}

	err := trc.topicReferenceService.UpdateTopicReference(
		ctx,
		service.UpdateTopicReferenceInput{
			Name:      input.Name,
			Path:      input.Path,
			ID:        input.ID,
			TeacherID: teacher.ID,
		},
	)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, nil)
}

// @Summary TeacherDeleteTopicReference
// @Description Teacher delete topic reference
// @Tags topic reference
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /topic-references/teachers/{id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (trc *TopicReferenceController) TeacherDeleteTopicReference(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	userContext := context_util.GetUserContext(ctx)
	if userContext == nil {
		response.ErrorResponse(ctx, http.StatusNotFound, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
		return
	}

	var teacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Select("id").Where("user_id = ?", userContext.ID).First(&teacher).Error; err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
		return
	}

	err = trc.topicReferenceService.DeleteTopicReference(
		ctx,
		service.DeleteTopicReferenceInput{
			ID:        int64(id),
			TeacherID: teacher.ID,
		},
	)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, nil)
}

// @Summary AdminCreateTopicReference
// @Description Admin create topic reference
// @Tags topic reference
// @Accept json
// @Produce json
// @Param data body topic_reference_dto.AdminCreateTopicReferenceInput true "data"
// @Router /topic-references/admins [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (trc *TopicReferenceController) AdminCreateTopicReference(ctx *gin.Context) {
	var input topic_reference_dto.AdminCreateTopicReferenceInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	err := trc.topicReferenceService.AdminCreateTopicReference(
		ctx,
		service.CreateTopicReferenceInput{
			Name:      input.Name,
			Path:      input.Path,
			TeacherID: input.TeacherID,
		},
	)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, nil)
}
