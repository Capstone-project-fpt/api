package jwt_util

import (
	"errors"
	"time"

	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type JwtInput struct {
	UserId int64
}

type ResetPassJwtInput struct {
	UserId int64
	Email  string
}

var jwtConfig = global.Config.Jwt

func GenerateAccessToken(payload JwtInput) (string, error) {
	secretKey := []byte(jwtConfig.Secret)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.UserId,
		"iss": global.Config.Server.Name,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(jwtConfig.Expiration) * time.Second).UnixMilli(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(payload JwtInput) (string, error) {
	secretKey := []byte(jwtConfig.RefreshSecret)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":           payload.UserId,
		"iss":           global.Config.Server.Name,
		"iat":           time.Now().Unix(),
		"exp":           time.Now().Add(time.Duration(jwtConfig.RefreshExpiration) * time.Second).UnixMilli(),
		"refresh_token": true,
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateResetPasswordToken(payload ResetPassJwtInput) (string, error) {
	secretKey := []byte(jwtConfig.Secret)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload,
		"iss": global.Config.Server.Name,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(constant.DefaultResetPasswordTokenExpiration) * time.Second).UnixMilli(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyTokenResetPassword(tokenString string) (*ResetPassJwtInput, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.TokenInvalid,
		})
		
		return nil, errors.New(message)
	}
	
	claims := token.Claims.(jwt.MapClaims)
	var payload ResetPassJwtInput
	payload.Email = claims["sub"].(map[string]interface{})["Email"].(string)

	return &payload, nil
}
