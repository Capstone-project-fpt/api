package service

import (
	"database/sql"
	"errors"
	"net/http"

	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/repository"
	"github.com/api/internal/types"
	"github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type IAuthService interface {
	Register(ctx *gin.Context, email, password string) (int, error)
	Login(ctx *gin.Context, email string, password string) (string, string, int, error)
}

type authService struct {
	userRepository repository.IUserRepository
	authProcessService IAuthProcessService
}

func NewAuthService(userRepository repository.IUserRepository, authProcessService IAuthProcessService) IAuthService {
	return &authService{
		userRepository: userRepository,
		authProcessService: authProcessService,
	}
}

func (as *authService) Register(ctx *gin.Context, email, password string) (int, error) {
	_, err := as.userRepository.GetUserByEmail(ctx, email)

	if err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserAlreadyExists,
		})

		return http.StatusConflict, errors.New(message)
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	err = as.userRepository.CreateUser(ctx, database.CreateUserParams{
		Email:    email,
		Password: sql.NullString{String: string(hashedPassword), Valid: true},
		Name:     "Admin FPT",
		UserType: constant.UserType.Admin,
	})

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	return http.StatusCreated, nil
}

func (as *authService) Login(ctx *gin.Context, email string, password string) (string, string, int, error) {
	user, err := as.userRepository.GetUserByEmail(ctx, email)

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return "", "", http.StatusNotFound, errors.New(message)
	}

	if !utils.CheckPasswordHash(password, user.Password.String) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return "", "", http.StatusNotFound, errors.New(message)
	}

	userContext := types.NewUserContext(user)

	accessToken, refreshToken, err := as.authProcessService.ResolveAccessAndRefreshToken(ctx, &userContext)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return "", "", http.StatusInternalServerError, errors.New(message)
	}

	return accessToken, refreshToken, http.StatusOK, nil
}
