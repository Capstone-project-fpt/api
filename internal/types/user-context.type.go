package types

import (
	"github.com/api/database/model"
)

type UserContext struct {
	ID       int64
	Name     string
	UserType string
	Email    string
}

func NewUserContext(user *model.User) UserContext {
	userContext := UserContext{
		ID:       user.ID,
		Name:     user.Name,
		UserType: user.UserType,
		Email:    user.Email,
	}

	return userContext
}
