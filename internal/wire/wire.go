//go:build wireinject

package wire

import (
	"github.com/api/internal/controller"
	capstone_group_controller "github.com/api/internal/controller/capstone_group"
	"github.com/api/internal/queue"
	"github.com/api/internal/service"
	admin_service "github.com/api/internal/service/admin"
	auth_service "github.com/api/internal/service/auth"
	capstone_group_service "github.com/api/internal/service/capstone_group"
	"github.com/api/pkg/service/aws"
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
		queue.NewEmailNewAccountsPublisher,
		admin_service.NewAdminService,
		service.NewUserService,
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

func InitializeUploadController() *controller.UploadController {
	wire.Build(
		aws.NewAwsS3Service,
		service.NewUploadService,
		controller.NewUploadController,
	)

	return &controller.UploadController{}
}

func InitializeSemesterController() *controller.SemesterController {
	wire.Build(
		service.NewSemesterService,
		controller.NewSemesterController,
	)

	return &controller.SemesterController{}
}

func InitializeCapstoneGroupController() *capstone_group_controller.CapstoneGroupController {
	wire.Build(
		queue.NewEmailInviteMentorToCapstoneGroupPublisher,
		capstone_group_service.NewCapstoneGroupService,
		capstone_group_service.NewCapstoneGroupTopicService,
		capstone_group_controller.NewCapstoneGroupController,
	)

	return &capstone_group_controller.CapstoneGroupController{}
}

func InitializeEvaluationCommitteeController() *controller.EvaluationCommitteeController {
	wire.Build(
		service.NewEvaluationCommitteeService,
		controller.NewEvaluationCommitteeController,
	)

	return &controller.EvaluationCommitteeController{}
}

func InitializeScheduleReviewController() *controller.ScheduleReviewController {
	wire.Build(
		service.NewScheduleReviewService,
		controller.NewScheduleReviewController,
	)

	return &controller.ScheduleReviewController{}
}