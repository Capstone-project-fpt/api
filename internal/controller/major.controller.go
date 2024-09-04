package controller

import (
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/major_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type MajorController struct {
	majorService service.IMajorService
}

func NewMajorController(majorService service.IMajorService) *MajorController {
	return &MajorController{
		majorService: majorService,
	}
}

// @Summary GetListMajor
// @Description Get list major
// @Tags Public
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Router /majors [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} major_dto.OutputGetListMajor
func (mc *MajorController) GetListMajor(ctx *gin.Context) {
	var input major_dto.InputGetListMajor
	localizer := global.Localizer
	if err := ctx.ShouldBindQuery(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	result, err := mc.majorService.GetListMajor(ctx, input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, result)
}
