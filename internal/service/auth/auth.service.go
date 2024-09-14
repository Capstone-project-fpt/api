package auth_service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/auth_dto"
	"github.com/api/internal/types"
	"github.com/api/pkg/mail"
	jwt_util "github.com/api/pkg/utils/jwt"
	password_util "github.com/api/pkg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

const (
	resetPasswordTokenKey = "reset_password_%s"
)

type IAuthService interface {
	Register(ctx *gin.Context, email, password string) (int, error)
	Login(ctx *gin.Context, email string, password string) (string, string, int, error)
	LoginGoogleHandle(ctx *gin.Context)
	LoginGoogleCallbackHandle(ctx *gin.Context) (string, error)
	ForgotPassword(ctx *gin.Context, email string) error
	ResetPassword(ctx *gin.Context, input *auth_dto.InputResetPassword) (int, error)
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

func (as *authService) ResetPassword(ctx *gin.Context, input *auth_dto.InputResetPassword) (int, error) {
	payload, err := jwt_util.VerifyTokenResetPassword(input.Token)

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return http.StatusBadRequest, errors.New(message)
	}

	var user model.User
	if err := global.Db.Model(model.User{}).Select("id").Where("email = ?", payload.Email).First(&user).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return http.StatusBadRequest, errors.New(message)
	}

	redis := global.RDb
	key := fmt.Sprintf(resetPasswordTokenKey, payload.Email)

	if _, err := redis.Get(ctx, key).Result(); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TokenInvalid,
		})

		return http.StatusBadRequest, errors.New(message)
	}

	if err := redis.Del(ctx, key).Err(); err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	hashPassword, err := password_util.HashPassword(input.Password)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	if err = global.Db.Model(model.User{}).Where("id = ?", user.ID).Update("password", hashPassword).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return http.StatusInternalServerError, errors.New(message)
	}

	return http.StatusOK, nil
}

func (as *authService) ForgotPassword(ctx *gin.Context, email string) error {
	isAlreadySend := as.checkAlreadySendResetPasswordLink(ctx, email)

	if isAlreadySend {
		// NOTE: Can return error if already send link to user
		return nil
	}

	redis := global.RDb
	var user model.User
	if err := global.Db.Model(model.User{}).Select("id", "email", "name").Where("email = ?", email).First(&user).Error; err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		})

		return errors.New(message)
	}

	token, err := jwt_util.GenerateResetPasswordToken(jwt_util.ResetPassJwtInput{Email: user.Email, UserId: user.ID})

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to generate reset password token, Error: ", zap.Error(err))

		return errors.New(message)
	}

	key := fmt.Sprintf(resetPasswordTokenKey, user.Email)
	redis.Set(ctx, key, token, time.Duration(constant.DefaultResetPasswordTokenExpiration)*time.Second)

	data := mail.MailResetPasswordTemplateData{
		Name:      user.Name,
		ResetLink: fmt.Sprintf("%s/auth/reset-password?token=%s", global.Config.Server.WebURL, token),
	}

	err = mail.SendResetPasswordEmail(user.Email, data)

	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to send email, Error: ", zap.Error(err))

		return errors.New(message)
	}

	err = as.clearTokenSessions(ctx, email)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to send email, Error: ", zap.Error(err))

		return errors.New(message)
	}

	return nil
}

func (as *authService) checkAlreadySendResetPasswordLink(ctx *gin.Context, email string) bool {
	redis := global.RDb
	key := fmt.Sprintf(resetPasswordTokenKey, email)
	_, err := redis.Get(ctx, key).Result()

	return err == nil
}

func (as *authService) clearTokenSessions(ctx *gin.Context, email string) error {
	redis := global.RDb
	var tokenKeys []string
	accessTokenIter := redis.Scan(ctx, 0, fmt.Sprintf("%s_%s*", constant.RedisKey.ActiveAccessToken, email), 0).Iterator()
	for accessTokenIter.Next(ctx) {
		tokenKeys = append(tokenKeys, accessTokenIter.Val())
	}
	if err := accessTokenIter.Err(); err != nil {
		return err
	}

	refreshTokenIter := redis.Scan(ctx, 0, fmt.Sprintf("%s_%s*", constant.RedisKey.ActiveRefreshToken, email), 0).Iterator()
	for refreshTokenIter.Next(ctx) {
		tokenKeys = append(tokenKeys, refreshTokenIter.Val())
	}
	if err := refreshTokenIter.Err(); err != nil {
		return err
	}

	var tokens []string

	for _, tokenKey := range tokenKeys {
		val, err := redis.Get(ctx, tokenKey).Result()
		if err != nil {
			return err
		}
		tokens = append(tokens, val)
	}

	var keys []string
	keys = append(keys, tokenKeys...)
	keys = append(keys, tokens...)

	if len(keys) > 0 {
		err := redis.Del(ctx, keys...).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
