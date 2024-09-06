package controller

import (
	"errors"
	"net/http"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/internal/service"
	"github.com/api/pkg/response"
	"github.com/api/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AdminController struct {
	adminService service.IAdminService
}

func NewAdminController(adminService service.IAdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// @Summary CreateStudentAccount
// @Description Admin Create Student Account
// @Tags Admin
// @Accept json
// @Produce json
// @Param data body admin_dto.InputAdminCreateStudentAccount true "data"
// @Success 200 {object} auth_dto.OutputLogin
// @Router /admin/students/create-account [post]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
func (ac *AdminController) CreateStudentAccount(ctx *gin.Context) {
	var input admin_dto.InputAdminCreateStudentAccount
	err := validateCreateStudentAccount(ctx, &input)

	if err != nil {
		return
	}

	statusCode, err := ac.adminService.CreateStudentAccount(ctx, &input)

	if err != nil {
		response.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

	response.SuccessResponse(ctx, 200, dto.OutputCommon{Message: ""})
}

func validateCreateStudentAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateStudentAccount) error {
	localizer := global.Localizer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidParams,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return err
	}

	if !validator.IsValidFptEmail(input.Email) {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidStudentEmailFPT,
		})

		response.ErrorResponse(ctx, http.StatusBadRequest, message)
		return errors.New(message)
	}

	return nil
}