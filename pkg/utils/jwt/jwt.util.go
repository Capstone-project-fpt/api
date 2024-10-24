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

type InviteMentorJwtInput struct {
	TeacherID       int64
	CapstoneGroupID int64
}


func GenerateAccessToken(payload JwtInput) (string, error) {
	secretKey := []byte(global.Config.Jwt.Secret)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.UserId,
		"iss": global.Config.Server.Name,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(global.Config.Jwt.Expiration) * time.Second).UnixMilli(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(payload JwtInput) (string, error) {
	secretKey := []byte(global.Config.Jwt.RefreshSecret)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":           payload.UserId,
		"iss":           global.Config.Server.Name,
		"iat":           time.Now().Unix(),
		"exp":           time.Now().Add(time.Duration(global.Config.Jwt.RefreshExpiration) * time.Second).UnixMilli(),
		"refresh_token": true,
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateResetPasswordToken(payload ResetPassJwtInput) (string, error) {
	secretKey := []byte(global.Config.Jwt.Secret)
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

func GenerateInviteMentorToken(payload InviteMentorJwtInput) (string, error) {
	secretKey := []byte(global.Config.Jwt.Secret)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload,
		"iss": global.Config.Server.Name,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(constant.DefaultInviteMentorTokenLength) * time.Second).UnixMilli(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyInviteMentorToken(tokenString string) (*InviteMentorJwtInput, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
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
	var payload InviteMentorJwtInput
	payload.TeacherID = int64(claims["sub"].(map[string]interface{})["TeacherID"].(float64))
	payload.CapstoneGroupID = int64(claims["sub"].(map[string]interface{})["CapstoneGroupID"].(float64))

	return &payload, nil
}

func VerifyTokenResetPassword(tokenString string) (*ResetPassJwtInput, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
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
