package service

import (
	"net/http"

	"github.com/go-ecommerce-backend-api/internal/repository"
)

type IUserService interface {
	Register(email string, password string) int
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (userService *userService) Register(email string, password string) int {
	if !userService.userRepository.GetUserByEmail(email) {
		return http.StatusBadRequest
	}
	
	return http.StatusAccepted
}