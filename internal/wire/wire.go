//go:build wireinject

package wire

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/repository"
	"github.com/api/internal/service"
	"github.com/google/wire"
)

func InitializeUserController() *controller.UserController {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return &controller.UserController{}
}

func InitializeAuthController() *controller.AuthController {
	wire.Build(
		repository.NewUserRepository,
		service.NewAuthService,
		controller.NewAuthController,
	)

	return &controller.AuthController{}
}