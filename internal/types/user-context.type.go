package types

import database "github.com/api/database/sqlc"

type UserContext struct {
	ID       int64
	Name     string
	UserType string
	Email    string
}

// type UserContext struct {
// 	ID       int64   `json:"id"`
// 	Name     string  `json:"name"`
// 	UserType string  `json:"user_type"`
// 	Email    string  `json:"email"`
// 	Code     *string `json:"code"`
// }

func NewUserContext(user database.User) UserContext {
	userContext := UserContext{
		ID:       user.ID,
		Name:     user.Name,
		UserType: user.UserType,
		Email:    user.Email,
	}

	return userContext
}
