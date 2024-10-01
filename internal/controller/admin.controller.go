package controller

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/admin_dto"
	admin_service "github.com/api/internal/service/admin"
	"github.com/api/pkg/response"
	file_util "github.com/api/pkg/utils/file"
	"github.com/api/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AdminController struct {
	adminService admin_service.IAdminService
}

func NewAdminController(adminService admin_service.IAdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// @Summary UploadFileStudentData
// @Description Admin upload Excel file to create student accounts
// @Tags Admin
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file"
// @Router /admin/students/import-data [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) UploadFileStudentData(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}
	err = validateExcelFile(file)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	statusCode, importOutput := ac.adminService.UploadFileStudentData(ctx, file)

	if statusCode != http.StatusOK {
		response.ErrorResponse(ctx, statusCode, importOutput)
		return
	}

	response.SuccessResponse(ctx, statusCode, importOutput)
}

// @Summary UploadFileTeacherData
// @Description Admin upload Excel file to create teacher accounts
// @Tags Admin
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file"
// @Router /admin/teachers/import-data [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) UploadFileTeacherData(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return
	}
	err = validateExcelFile(file)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	statusCode, importOutput := ac.adminService.UploadFileTeacherData(ctx, file)

	if statusCode != http.StatusOK {
		response.ErrorResponse(ctx, statusCode, importOutput)
		return
	}

	response.SuccessResponse(ctx, statusCode, importOutput)
}

// @Summary CreateStudentAccount
// @Description Admin Create Student Account
// @Tags Admin
// @Accept json
// @Produce json
// @Param data body admin_dto.InputAdminCreateStudentAccount true "data"
// @Router /admin/students/create-account [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) CreateStudentAccount(ctx *gin.Context) {
	var input admin_dto.InputAdminCreateStudentAccount
	err := validateCreateAccount(ctx, &input)

	if err != nil {
		return
	}

	statusCode, err := ac.adminService.CreateStudentAccount(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

	localizer := global.Localizer
	message := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateStudentAccountSuccess,
	})

	response.SuccessResponse(ctx, statusCode, dto.OutputCommon{Message: message})
}

// @Summary CreateTeacherAccount
// @Description Admin Create Teacher Account
// @Tags Admin
// @Accept json
// @Produce json
// @Param data body admin_dto.InputAdminCreateTeacherAccount true "data"
// @Router /admin/teachers/create-account [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) CreateTeacherAccount(ctx *gin.Context) {
	var input admin_dto.InputAdminCreateTeacherAccount
	err := validateCreateAccount(ctx, &input)

	if err != nil {
		return
	}

	statusCode, err := ac.adminService.CreateTeacherAccount(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

	localizer := global.Localizer
	message := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.CreateTeacherAccountSuccess,
	})

	response.SuccessResponse(ctx, statusCode, dto.OutputCommon{Message: message})
}

func validateExcelFile(file *multipart.FileHeader) error {
	isExcelFile := file_util.IsExcelFile(file)

	if !isExcelFile {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		return errors.New(message)
	}

	return nil
}

func validateCreateAccount(ctx *gin.Context, input admin_dto.AccountWithEmail) error {
	localizer := global.Localizer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return err
	}

	if !validator.IsValidFptEmail(input.GetEmail()) {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidStudentEmailFPT,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return errors.New(message)
	}

	return nil
}
