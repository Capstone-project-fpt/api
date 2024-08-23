package repository

import (
	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	GetUserByEmail(ctx *gin.Context, email string) (database.User, error)
	GetUserById(ctx *gin.Context, id int64) (database.User, error)
}

type userRepository struct {}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUserByEmail(ctx *gin.Context, email string) (database.User, error) {
	user, err := global.Db.GetUserByEmail(ctx, email)

	return user, err
}

func (u *userRepository) GetUserById(ctx *gin.Context, id int64) (database.User, error) {
	user, err := global.Db.GetUserById(ctx, id)

	return user, err
}