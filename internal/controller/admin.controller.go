package controller

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/internal/service"
	admin_service "github.com/api/internal/service/admin"
	"github.com/api/pkg/response"
	util "github.com/api/pkg/utils"
	file_util "github.com/api/pkg/utils/file"
	"github.com/api/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AdminController struct {
	adminService admin_service.IAdminService
	userService  service.IUserService
}

func NewAdminController(adminService admin_service.IAdminService, userService service.IUserService) *AdminController {
	return &AdminController{
		adminService: adminService,
		userService:  userService,
	}
}

// @Summary GetUser
// @Description Get User
// @Tags Admin
// @Produce json
// @Param id path int true "id"
// @Router /admin/users/{id} [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} user_dto.GetUserSwaggerOutput
// @Security ApiKeyAuth
func (ac *AdminController) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	outputGetUser, err := ac.userService.GetUser(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, outputGetUser)
}

// @Summary DeleteUser
// @Description Delete a user by ID
// @Tags Admin
// @Produce json
// @Param id path int true "id"
// @Router /admin/users/{id} [delete]
// @Failure 400 {object} response.ResponseErr
// @Failure 404 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ac.userService.DeleteUser(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, dto.OutputCommon{Message: "User deleted successfully"})
}

// @Summary GetListUsers
// @Description Get list user
// @Tags Admin
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param user_types query []string false "UserTypes" collectionFormat(multi)
// @Param email query string false "Email"
// @Router /admin/users [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} admin_dto.ListUsersOutput
// @Security ApiKeyAuth
func (ac *AdminController) GetListUsers(ctx *gin.Context) {
	var input admin_dto.GetListUsersInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.Offset, _ = util.GetPagination(int(input.Page), int(input.Limit))
	result, err := ac.userService.GetListUsers(ctx, service.GetListUsersInput{
		Limit:     input.Limit,
		Page:      input.Page,
		Offset:    input.Offset,
		UserTypes: input.UserTypes,
		Email:     input.Email,
	})

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, result)
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
// @Param data body admin_dto.AdminCreateStudentAccountInput true "data"
// @Router /admin/students/create-account [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) CreateStudentAccount(ctx *gin.Context) {
	var input admin_dto.AdminCreateStudentAccountInput
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
// @Param data body admin_dto.AdminCreateTeacherAccountInput true "data"
// @Router /admin/teachers/create-account [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (ac *AdminController) CreateTeacherAccount(ctx *gin.Context) {
	var input admin_dto.AdminCreateTeacherAccountInput
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
