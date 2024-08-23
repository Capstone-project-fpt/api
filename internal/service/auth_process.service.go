package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/types"
	jwt_util "github.com/api/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

type IAuthProcessService interface {
	ResolveAccessAndRefreshToken (ctx *gin.Context, userContext *types.UserContext) (string, string, error)
}

type authProcessService struct {}

func NewAuthProcessService() IAuthProcessService {
	return &authProcessService{}
}

func (s *authProcessService) ResolveAccessAndRefreshToken(ctx *gin.Context, userContext *types.UserContext) (string, string, error) {
	accessToken, refreshToken, err := generateTokens(userContext)
	if err != nil {
		return "", "", err
	}

	err = storeTokensInRedis(ctx, accessToken, refreshToken, userContext)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateTokens(userContext *types.UserContext) (string, string, error) {
	accessTokenChan := make(chan string)
	refreshTokenChan := make(chan string)
	errChan := make(chan error)

	go generateToken(accessTokenChan, errChan, func() (string, error) {
		return jwt_util.GenerateAccessToken(jwt_util.JwtInput{UserId: userContext.ID})
	})

	go generateToken(refreshTokenChan, errChan, func() (string, error) {
		return jwt_util.GenerateRefreshToken(jwt_util.JwtInput{UserId: userContext.ID})
	})

	var accessToken, refreshToken string
	for i := 0; i < 2; i++ {
		select {
		case token := <-accessTokenChan:
			accessToken = token
		case token := <-refreshTokenChan:
			refreshToken = token
		case err := <-errChan:
			return "", "", err
		}
	}

	return accessToken, refreshToken, nil
}

func generateToken(tokenChan chan string, errChan chan error, tokenFunc func() (string, error)) {
	token, err := tokenFunc()
	if err != nil {
		errChan <- err
		return
	}
	tokenChan <- token
}

func storeTokensInRedis(ctx *gin.Context, accessToken, refreshToken string, userContext *types.UserContext) error {
	jwtConfig := global.Config.Jwt

	userContextJson, err := json.Marshal(userContext)
	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()
	errChan := make(chan error)

	go storeTokenInRedis(ctx, errChan, accessToken, userContextJson, constant.RedisKey.ActiveAccessToken, timestamp, jwtConfig.Expiration)
	go storeTokenInRedis(ctx, errChan, refreshToken, userContextJson, constant.RedisKey.ActiveRefreshToken, timestamp, jwtConfig.RefreshExpiration)

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

func storeTokenInRedis(ctx *gin.Context, errChan chan error, token string, userContextJson []byte, keyPrefix string, timestamp int64, expiration int) {
	redis := global.RDb

	_, err := redis.Set(ctx, token, userContextJson, time.Duration(expiration)*time.Second).Result()
	if err != nil {
		errChan <- err
		return
	}

	activeTokenKey := fmt.Sprintf("%s_%d_%d", keyPrefix, timestamp, timestamp)
	_, err = redis.Set(ctx, activeTokenKey, token, time.Duration(expiration)*time.Second).Result()
	if err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
