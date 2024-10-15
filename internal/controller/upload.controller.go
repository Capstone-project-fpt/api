package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/upload_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UploadController struct {
	uploadService service.IUploadService
}

func NewUploadController(uploadService service.IUploadService) *UploadController {
	return &UploadController{
		uploadService: uploadService,
	}
}

// @Summary GenerateUploadPresignUrl
// @Description Generate Upload PresignUrl
// @Tags Upload
// @Accept json
// @Produce json
// @Param data body upload_dto.GenerateUploadPresignUrlInput true "data"
// @Router /upload/presign-url [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (uc *UploadController) GenerateUploadPresignUrl(ctx *gin.Context) {
	var input upload_dto.GenerateUploadPresignUrlInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	preSignUrl, err := uc.uploadService.GenerateUploadPresignUrl(input.Key)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, preSignUrl)
}
