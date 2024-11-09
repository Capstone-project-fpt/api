package router

import (
	"github.com/api/internal/router/admin"
	capstone_group_router "github.com/api/internal/router/capstone_group"
	"github.com/api/internal/router/evaluation_committee"
	"github.com/api/internal/router/public"
	"github.com/api/internal/router/schedule_review"
	"github.com/api/internal/router/semester"
	"github.com/api/internal/router/topic_reference"
	"github.com/api/internal/router/upload"
	"github.com/api/internal/router/user"
)

type RouterGroup struct {
	Public              public.PublicRouterGroup
	User                user.UserRouterGroup
	Admin               admin.AdminRouterGroup
	TopicReference      topic_reference.TopicReferenceGroup
	Upload              upload.UploadGroup
	Semester            semester.SemesterRouterGroup
	CapstoneGroup       capstone_group_router.CapstoneGroupRouterGroup
	EvaluationCommittee evaluation_committee.EvaluationCommitteeGroup
	ScheduleReview      schedule_review.ScheduleReviewGroup
}

var RouterGroupApp = new(RouterGroup)
