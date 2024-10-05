//go:build wireinject

package wire

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/service"
	auth_service "github.com/api/internal/service/auth"
	admin_service "github.com/api/internal/service/admin"
	"github.com/google/wire"
	"github.com/api/internal/queue"
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
		queue.NewEmailNewAccountsPublisher,
		admin_service.NewAdminService,
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

func InitializeSubMajorController() *controller.SubMajorController {
	wire.Build(
		service.NewSubMajorService,
		controller.NewSubMajorController,
	)

	return &controller.SubMajorController{}
}

func InitializeTopicReferenceController() *controller.TopicReferenceController {
	wire.Build(
		service.NewTopicReferenceService,
		controller.NewTopicReferenceController,
	)

	return &controller.TopicReferenceController{}
}