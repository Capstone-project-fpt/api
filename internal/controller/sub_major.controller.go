package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/sub_major_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SubMajorController struct {
	subMajorService service.ISubMajorService
}

func NewSubMajorController(subMajorService service.ISubMajorService) *SubMajorController {
	return &SubMajorController{
		subMajorService: subMajorService,
	}
}

// @Summary GetListSubMajor
// @Description Get list sub major
// @Tags Public
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param major_id query int false "Major ID"
// @Router /sub-majors [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} sub_major_dto.OutputGetListMajor
func (mc *SubMajorController) GetListSubMajor(ctx *gin.Context) {
	var input sub_major_dto.InputGetListSubMajor
	localizer := global.Localizer
	if err := ctx.ShouldBindQuery(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	result, err := mc.subMajorService.GetListSubMajor(ctx, input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, result)
}
