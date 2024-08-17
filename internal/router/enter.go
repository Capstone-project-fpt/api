package router

import (
	"github.com/go-ecommerce-backend-api/internal/router/manager"
	"github.com/go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)