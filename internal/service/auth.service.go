package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/repository"
	"github.com/api/internal/types"
	"github.com/api/pkg/utils"
	jwt_util "github.com/api/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

type IAuthService interface {
	Register(ctx *gin.Context, email, password string) (int, error)
	Login(ctx *gin.Context, email string, password string) (string, string, int, error)
}

type authService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) IAuthService {
	return &authService{
		userRepository: userRepository,
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
		Name: "Admin FPT",
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

	accessToken, refreshToken, err := resolveAccessAndRefreshToken(ctx, userContext)
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})

		return "", "", http.StatusInternalServerError, errors.New(message)
	}

	return accessToken, refreshToken, http.StatusOK, nil
}

func resolveAccessAndRefreshToken(ctx *gin.Context, userContext types.UserContext) (string, string, error) {
	redis := global.RDb
	jwtConfig := global.Config.Jwt

	accessToken, err := jwt_util.GenerateAccessToken(jwt_util.JwtInput{UserId: userContext.ID})
	if err != nil {
		global.Logger.Info("Failed to generate access token", zap.Error(err))

		return "", "", err
	}

	refreshToken, err := jwt_util.GenerateRefreshToken(jwt_util.JwtInput{UserId: userContext.ID})
	if err != nil {
		global.Logger.Info("Failed to generate refresh token", zap.Error(err))

		return "", "", err
	}

	userContextJson, err := json.Marshal(userContext)

	if err != nil {
		global.Logger.Info("Failed to marshal user context", zap.Error(err))

		return "", "", err
	}

	_, err = redis.Set(ctx, accessToken, userContextJson, time.Duration(jwtConfig.Expiration) * time.Second).Result()
	if err != nil {
		global.Logger.Info("Failed to set access token to redis", zap.Error(err))

		return "", "", err
	}

	_, err = redis.Set(ctx, refreshToken, userContextJson, time.Duration(jwtConfig.RefreshExpiration) * time.Second).Result()
	if err != nil {
		global.Logger.Info("Failed to set refresh token to redis", zap.Error(err))

		return "", "", err
	}

	timestamp := time.Now().Unix()
	activeAccessToken := fmt.Sprintf("%s_%d_%d", constant.RedisKey.ActiveAccessToken, userContext.ID, timestamp)
	activeRefreshToken := fmt.Sprintf("%s_%d_%d", constant.RedisKey.ActiveRefreshToken, userContext.ID, timestamp)

	_, err = redis.Set(ctx, activeAccessToken, accessToken, time.Duration(jwtConfig.Expiration) * time.Second).Result()
	if err != nil {
		global.Logger.Info("Failed to set active access token to redis", zap.Error(err))

		return "", "", err
	}

	_, err = redis.Set(ctx, activeRefreshToken, refreshToken, time.Duration(jwtConfig.RefreshExpiration) * time.Second).Result()
	if err != nil {
		global.Logger.Info("Failed to set active refresh token to redis", zap.Error(err))

		return "", "", err
	}

	return accessToken, refreshToken, nil
}
