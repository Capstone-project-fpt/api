package auth_service

import (
	"errors"
	"net/http"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/types"
	password_util "github.com/api/pkg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type IAuthService interface {
	Register(ctx *gin.Context, email, password string) (int, error)
	Login(ctx *gin.Context, email string, password string) (string, string, int, error)
}

type authService struct {
	authProcessService IAuthProcessService
}

func NewAuthService(authProcessService IAuthProcessService) IAuthService {
	return &authService{
		authProcessService: authProcessService,
	}
}

// Note: this function it just used to create Admin account in dev, for production, will not support this function
func (as *authService) Register(ctx *gin.Context, email, password string) (int, error) {
	var user model.User
	err := global.Db.Model(model.User{}).Select("id").First(&user, "email = ?", email).Error

	if err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserAlreadyExists,
		})

		return http.StatusConflict, errors.New(message)
	}

	hashedPassword, err := password_util.HashPassword(password)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	newUser := model.User{
		Email:       email,
		Password:    hashedPassword,
		Name:        "Admin FPT",
		UserType:    constant.UserType.Admin,
		PhoneNumber: "0914121791",
	}
	err = global.Db.Create(&newUser).Error

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	return http.StatusCreated, nil
}

func (as *authService) Login(ctx *gin.Context, email string, password string) (string, string, int, error) {
	var user model.User
	err := global.Db.Model(model.User{}).Select("id", "email", "password", "user_type", "name").Find(&user, "email = ?", email).Error

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return "", "", http.StatusNotFound, errors.New(message)
	}

	if !password_util.CheckPasswordHash(password, user.Password) {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return "", "", http.StatusNotFound, errors.New(message)
	}

	userContext := types.NewUserContext(&user)

	accessToken, refreshToken, err := as.authProcessService.ResolveAccessAndRefreshToken(ctx, &userContext)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return "", "", http.StatusInternalServerError, errors.New(message)
	}

	return accessToken, refreshToken, http.StatusOK, nil
}
