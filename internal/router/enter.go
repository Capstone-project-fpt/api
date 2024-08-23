package router

import (
	"github.com/api/internal/router/admin"
	"github.com/api/internal/router/public"
	"github.com/api/internal/router/user"
)

type RouterGroup struct {
	Public  public.PublicRouterGroup
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
