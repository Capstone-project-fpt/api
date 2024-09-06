//go:build wireinject

package wire

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/service"
	auth_service "github.com/api/internal/service/auth"
	"github.com/google/wire"
)

func InitializeUserController() *controller.UserController {
	wire.Build(
		service.NewUserService,
		controller.NewUserController,
	)

	return &controller.UserController{}
}

func InitializeAuthController() *controller.AuthController {
	wire.Build(
		auth_service.NewAuthProcessService,
		auth_service.NewAuthService,
		controller.NewAuthController,
	)

	return &controller.AuthController{}
}

func InitializeAdminController() *controller.AdminController {
	wire.Build(
		service.NewAdminService,
		controller.NewAdminController,
	)

	return &controller.AdminController{}
}

func InitializeMajorController() *controller.MajorController {
	wire.Build(
		service.NewMajorService,
		controller.NewMajorController,
	)

	return &controller.MajorController{}
}