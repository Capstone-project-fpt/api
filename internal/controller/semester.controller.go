package controller

import (
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/semester_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SemesterController struct {
	semesterService service.ISemesterService
}

func NewSemesterController(semesterService service.ISemesterService) *SemesterController {
	return &SemesterController{
		semesterService: semesterService,
	}
}

// @Summary GetListSemesters
// @Description Get list semester
// @Tags Semester
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Router /semesters [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} semester_dto.ListSemestersOutput
// @Security ApiKeyAuth
func (sc *SemesterController) GetListSemesters(ctx *gin.Context) {
	var input semester_dto.GetListSemestersInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	result, err := sc.semesterService.GetListSemester(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, result)
}

// @Summary GetSemester
// @Description Get Semester
// @Tags Semester
// @Produce json
// @Param id path int true "id"
// @Router /semesters/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} semester_dto.GetSemesterSwaggerOutput
// @Security ApiKeyAuth
func (sc *SemesterController) GetSemester(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	output, err := sc.semesterService.GetSemester(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, output)
}

// @Summary AdminCreateSemester
// @Description Admin Create Semester
// @Tags Semester
// @Accept json
// @Produce json
// @Param data body semester_dto.CreateSemesterInput true "data"
// @Router /semesters [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (sc *SemesterController) AdminCreateSemester(ctx *gin.Context) {
	var input semester_dto.CreateSemesterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err := sc.semesterService.CreateSemester(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateSemesterSuccess,
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
}

// @Summary AdminUpdateSemester
// @Description Admin Update Semester
// @Tags Semester
// @Accept json
// @Produce json
// @Param data body semester_dto.UpdateSemesterInput true "data"
// @Router /semesters [put]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (sc *SemesterController) AdminUpdateSemester(ctx *gin.Context) {
	var input semester_dto.UpdateSemesterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err := sc.semesterService.UpdateSemester(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateSemesterSuccess,
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
}

// @Summary AdminDeleteSemester
// @Description Admin Delete Semester
// @Tags Semester
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /semesters/{id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (sc *SemesterController) AdminDeleteSemester(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = sc.semesterService.DeleteSemester(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.DeleteSemesterSuccess,
	})

	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: message})
}