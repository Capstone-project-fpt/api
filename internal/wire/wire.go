//go:build wireinject 

package wire

import (
	"github.com/go-ecommerce-backend-api/internal/controller"
	"github.com/go-ecommerce-backend-api/internal/repository"
	"github.com/go-ecommerce-backend-api/internal/service"
	"github.com/google/wire"
)

func InitializeUserController() (*controller.UserController, error) {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return &controller.UserController{}, nil
}
