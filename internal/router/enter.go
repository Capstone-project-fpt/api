package router

import (
	"github.com/api/internal/router/admin"
	"github.com/api/internal/router/public"
	"github.com/api/internal/router/topic_reference"
	"github.com/api/internal/router/user"
)

type RouterGroup struct {
	Public         public.PublicRouterGroup
	User           user.UserRouterGroup
	Admin          admin.AdminRouterGroup
	TopicReference topic_reference.TopicReferenceGroup
}

var RouterGroupApp = new(RouterGroup)
