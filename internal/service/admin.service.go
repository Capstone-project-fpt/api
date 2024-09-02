package service

import (
	"database/sql"
	"errors"
	"net/http"

	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/internal/repository"
	"github.com/api/pkg/mail"
	password_util "github.com/api/pkg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

type IAdminService interface {
	CreateStudentAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateStudentAccount) (int, error)
}

type adminService struct {
	userRepository    repository.IUserRepository
	studentRepository repository.IStudentRepository
}

func NewAdminService(userRepository repository.IUserRepository, studentRepository repository.IStudentRepository) IAdminService {
	return &adminService{
		userRepository:    userRepository,
		studentRepository: studentRepository,
	}
}

func (as *adminService) CreateStudentAccount(ctx *gin.Context, input *admin_dto.InputAdminCreateStudentAccount) (int, error) {
	_, err := as.userRepository.GetUserByEmail(ctx, input.Email)

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

	tx, err := global.RawDb.Begin()
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to begin a transaction, Error: ", zap.Error(err))

		return http.StatusInternalServerError, errors.New(message)
	}
	defer tx.Rollback()
	qtx := global.Db.WithTx(tx)

	id, err := qtx.CreateUserAndReturnId(ctx, database.CreateUserAndReturnIdParams{
		Name:        input.Name,
		Password:    sql.NullString{String: string(hashPassword), Valid: true},
		Email:       input.Email,
		UserType:    constant.UserType.Student,
		PhoneNumber: input.PhoneNumber,
	})

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	err = qtx.CreateStudent(ctx, database.CreateStudentParams{
		Code:       input.Code,
		SubMajorID: input.SubMajorId,
		UserID:     id,
	})

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = tx.Commit()
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to commit a transaction, Error: ", zap.Error(err))

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
