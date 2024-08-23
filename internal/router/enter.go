package router

import (
	"github.com/api/internal/router/admin"
	"github.com/api/internal/router/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
