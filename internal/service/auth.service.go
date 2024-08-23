package service

import (
	"net/http"

	"github.com/api/internal/repository"
	"github.com/api/pkg/response"
	"github.com/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type IAuthService interface{}

type authService struct{
	userRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (as *authService) Login(ctx *gin.Context, email string, password string) (interface{}, *response.ResponseErr) {
	user, err := as.userRepository.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, &response.ResponseErr{
			Code: http.StatusNotFound,
			Success: false,
			Error: "User not found",
		}
	}

	if !utils.CheckPasswordHash(password, user.Password.String) {
		return nil, &response.ResponseErr{
			Code: http.StatusUnauthorized,
			Success: false,
			Error: "User not found",
		}
	}

	// validate password

	return user, nil
}
