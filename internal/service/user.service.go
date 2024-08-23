package service

import (
	"net/http"

	"github.com/api/internal/repository"
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
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

	if err != nil {
		return nil, &response.ResponseErr{
			Code: http.StatusNotFound,
			Success: false,
			Error: "Email not found",
		}
	}

	return http.StatusAccepted, nil
}
