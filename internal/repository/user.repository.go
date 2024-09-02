package repository

import (
	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/gin-gonic/gin"
)

type CreateUserParams database.CreateUserParams
type CreateUserAndReturnIdParams database.CreateUserAndReturnIdParams

type IUserRepository interface {
	CreateUser(ctx *gin.Context, arg CreateUserParams) error
	GetUserByEmail(ctx *gin.Context, email string) (database.User, error)
	GetUserById(ctx *gin.Context, id int64) (database.User, error)
	CreateUserAndReturnId(ctx *gin.Context, arg CreateUserAndReturnIdParams) (int64, error)
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (u *userRepository) CreateUser(ctx *gin.Context, arg CreateUserParams) error {
	return global.Db.CreateUser(ctx, database.CreateUserParams{
		Email:           arg.Email,
		Password:        arg.Password,
		Name:            arg.Name,
		UserType:        arg.UserType,
		PhoneNumber:     arg.PhoneNumber,
	})
}

func (u *userRepository) CreateUserAndReturnId(ctx *gin.Context, arg CreateUserAndReturnIdParams) (int64, error) {
	return global.Db.CreateUserAndReturnId(ctx, database.CreateUserAndReturnIdParams{
		Email:           arg.Email,
		Password:        arg.Password,
		Name:            arg.Name,
		UserType:        arg.UserType,
		PhoneNumber:     arg.PhoneNumber,
	})
}

func (u *userRepository) GetUserByEmail(ctx *gin.Context, email string) (database.User, error) {
	user, err := global.Db.GetUserByEmail(ctx, email)

	return user, err
}

func (u *userRepository) GetUserById(ctx *gin.Context, id int64) (database.User, error) {
	user, err := global.Db.GetUserById(ctx, id)

	return user, err
}
