package service

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/internal/dto/import_dto"
	"github.com/api/internal/queue"
	"github.com/api/internal/service"
	password_util "github.com/api/pkg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAdminService interface {
	CreateStudentAccount(ctx *gin.Context, input *admin_dto.AdminCreateStudentAccountInput) (int, error)
	CreateTeacherAccount(ctx *gin.Context, input *admin_dto.AdminCreateTeacherAccountInput) (int, error)
	UploadFileStudentData(ctx *gin.Context, file *multipart.FileHeader) (int, *import_dto.ImportOutput)
	UploadFileTeacherData(ctx *gin.Context, file *multipart.FileHeader) (int, *import_dto.ImportOutput)
	UpdateAccount(ctx *gin.Context, input *admin_dto.UpdateAccountInput) (int, error)
	AssignVerifierTopic(ctx *gin.Context, input *admin_dto.AdminAssignVerifierTopicInput) error
	UnassignVerifierTopic(ctx *gin.Context, input *admin_dto.AdminUnassignVerifierTopicInput) error
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

type adminService struct {
	emailNewAccountsPublisher queue.IBasePublisher[queue.EmailNewAccountsMessage]
	userService               service.IUserService
}

func NewAdminService(
	emailNewAccountsPublisher queue.IBasePublisher[queue.EmailNewAccountsMessage],
	userService service.IUserService,
) IAdminService {
	return &adminService{
		emailNewAccountsPublisher,
		userService,
	}
}

func (as *adminService) CreateStudentAccount(ctx *gin.Context, input *admin_dto.AdminCreateStudentAccountInput) (int, error) {
	return as.createAccount(ctx, InputCreateAccount{
		Name:        input.Name,
		Email:       input.Email,
		UserType:    constant.UserType.Student,
		PhoneNumber: input.PhoneNumber,
		SubMajorID:  input.SubMajorID,
		RoleType:    constant.RoleType.Student,
		Code:        input.Code,
	})
}

func (as *adminService) CreateTeacherAccount(ctx *gin.Context, input *admin_dto.AdminCreateTeacherAccountInput) (int, error) {
	return as.createAccount(ctx, InputCreateAccount{
		Name:        input.Name,
		Email:       input.Email,
		UserType:    constant.UserType.Teacher,
		PhoneNumber: input.PhoneNumber,
		SubMajorID:  input.SubMajorID,
		RoleType:    constant.RoleType.Teacher,
	})
}

func (as *adminService) createAccount(ctx *gin.Context, input InputCreateAccount) (int, error) {
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
			UserID:     user.ID,
			SubMajorID: input.SubMajorID,
			Code:       input.Code,
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

	var newAccounts []queue.NewAccountMessage
	newAccounts = append(newAccounts, queue.NewAccountMessage{
		Email:    input.Email,
		Name:     input.Name,
		Password: password,
	})

	if err := as.sendEmailNewAccountsCreated(ctx, newAccounts); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to send email, Error: ", zap.Error(err))

		return http.StatusInternalServerError, errors.New(message)
	}

	return http.StatusOK, nil
}

func (as *adminService) UpdateAccount(ctx *gin.Context, input *admin_dto.UpdateAccountInput) (int, error) {
	var user model.User
	err := global.Db.First(&user, input.UserID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.UserNotFound,
			})
			return http.StatusNotFound, errors.New(message)
		}
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.PhoneNumber != "" {
		user.PhoneNumber = input.PhoneNumber
	}

	tx := global.Db.Begin()
	if tx.Error != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UpdateFailed,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	switch user.UserType {
	case constant.UserType.Student:
		var student model.Student
		err := tx.First(&student, "user_id = ?", user.ID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: constant.MessageI18nId.StudentNotFound,
				})
				return http.StatusInternalServerError, errors.New(message)
			}
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			return http.StatusInternalServerError, errors.New(message)
		}

		if input.SubMajorID != 0 {
			student.SubMajorID = input.SubMajorID
		}
		if input.Code != "" {
			student.Code = input.Code
		}
		if err := tx.Save(&student).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.UpdateStudentDetailsFailed,
			})
			return http.StatusInternalServerError, errors.New(message)
		}

	case constant.UserType.Teacher:
		var teacher model.Teacher
		err := tx.First(&teacher, "user_id = ?", user.ID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: constant.MessageI18nId.TeacherNotFound,
				})
				return http.StatusInternalServerError, errors.New(message)
			}
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			return http.StatusInternalServerError, errors.New(message)
		}

		if input.SubMajorID != 0 {
			teacher.SubMajorID = input.SubMajorID
		}
		if err := tx.Save(&teacher).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.UpdateTeacherDetailsFailed,
			})
			return http.StatusInternalServerError, errors.New(message)
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return http.StatusInternalServerError, errors.New(message)
	}

	successMessage := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UpdateAccountSuccess,
	})
	return http.StatusOK, errors.New(successMessage)
}
