package service

import (
	"net/http"

	"github.com/api/internal/repository"
	"github.com/api/internal/constant"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type IUserService interface {
	Register(ctx *gin.Context, email string, password string) (interface{}, *response.ResponseErr)
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) Register(ctx *gin.Context, email string, password string) (interface{}, *response.ResponseErr) {
	_, err := us.userRepository.GetUserByEmail(ctx, email)
	loc, _ := ctx.Get("localizer")
	localizer := loc.(*i18n.Localizer)

	if err != nil {
		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.EmailNotFound,
		})

		return nil, &response.ResponseErr{
			Code:    http.StatusNotFound,
			Success: false,
			Error:   message,
		}
	}

	return http.StatusAccepted, nil
}
