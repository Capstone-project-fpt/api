package router

import (
	"github.com/api/internal/router/admin"
	"github.com/api/internal/router/public"
	"github.com/api/internal/router/topic_reference"
	"github.com/api/internal/router/upload"
	"github.com/api/internal/router/user"
)

type RouterGroup struct {
	Public         public.PublicRouterGroup
	User           user.UserRouterGroup
	Admin          admin.AdminRouterGroup
	TopicReference topic_reference.TopicReferenceGroup
	Upload         upload.UploadGroup
}

var RouterGroupApp = new(RouterGroup)
