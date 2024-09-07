package service

import (
	"errors"
	"net/http"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/pkg/mail"
	password_util "github.com/api/pkg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

type IAdminService interface {
	CreateStudentAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateStudentAccount) (int, error)
	CreateTeacherAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateTeacherAccount) (int, error)
}

type InputCreateAccount struct {
	Name        string
	Email       string
	UserType    string
	PhoneNumber string
	SubMajorID  int64
	RoleType    string
	Code        string
}

type adminService struct{}

func NewAdminService() IAdminService {
	return &adminService{}
}

func (as *adminService) CreateStudentAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateStudentAccount) (int, error) {
	return as.createAccount(InputCreateAccount{
		Name:        input.Name,
		Email:       input.Email,
		UserType:    constant.UserType.Student,
		PhoneNumber: input.PhoneNumber,
		SubMajorID:  input.SubMajorID,
		RoleType:    constant.RoleType.Student,
		Code:        input.Code,
	})
}

func (as *adminService) CreateTeacherAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateTeacherAccount) (int, error) {
	return as.createAccount(InputCreateAccount{
		Name:        input.Name,
		Email:       input.Email,
		UserType:    constant.UserType.Teacher,
		PhoneNumber: input.PhoneNumber,
		SubMajorID:  input.SubMajorID,
		RoleType:    constant.RoleType.Teacher,
	})
}

func (as *adminService) createAccount(input InputCreateAccount) (int, error) {
	var findUser model.User
	err := global.Db.Model(model.User{}).Select("id").First(&findUser, "email = ?", input.Email).Error

	if err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserAlreadyExists,
		})

		return http.StatusConflict, errors.New(message)
	}

	password := password_util.GenerateRandomPassword(constant.DefaultPasswordLength)
	hashPassword, err := password_util.HashPassword(password)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}
	tx := global.Db.Begin()
	if tx.Error != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to begin a transaction, Error: ", zap.Error(tx.Error))
		return http.StatusInternalServerError, errors.New(message)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	user := model.User{
		Name:        input.Name,
		Password:    hashPassword,
		Email:       input.Email,
		UserType:    input.UserType,
		PhoneNumber: input.PhoneNumber,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	switch input.UserType {
	case constant.UserType.Teacher:
		teacher := model.Teacher{
			SubMajorID: input.SubMajorID,
			UserID:     user.ID,
		}
		if err := tx.Create(&teacher).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, err
		}
	case constant.UserType.Student:
		student := model.Student{
			UserID: user.ID,
			SubMajorID: input.SubMajorID,
			Code: input.Code,
		}
		if err := tx.Create(&student).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, err
		}
	}
	var role model.Role
	if err = tx.Model(&model.Role{}).Select("id").First(&role, "name = ?", input.RoleType).Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	if err = tx.Exec("INSERT INTO users_roles (user_id, role_id) VALUES (?, ?)", user.ID, role.ID).Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	data := mail.MailNewAccountTemplateData{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	}

	err = mail.SendNewAccountEmail(input.Email, data)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to send email, Error: ", zap.Error(err))

		return http.StatusInternalServerError, errors.New(message)
	}

	return http.StatusOK, nil
}
