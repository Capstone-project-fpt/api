//go:build wireinject

package wire

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/repository"
	"github.com/api/internal/service"
	auth_service "github.com/api/internal/service/auth"
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
		auth_service.NewAuthProcessService,
		auth_service.NewAuthService,
		controller.NewAuthController,
	)

	return &controller.AuthController{}
}

func InitializeAdminController() *controller.AdminController {
	wire.Build(
		repository.NewUserRepository,
		repository.NewStudentRepository,
		service.NewAdminService,
		controller.NewAdminController,
	)

	return &controller.AdminController{}
}
