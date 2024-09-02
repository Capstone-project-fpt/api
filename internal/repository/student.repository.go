package repository

import (
	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/gin-gonic/gin"
)

type CreateStudentParams database.CreateStudentParams

type IStudentRepository interface {
	CreateStudent(ctx *gin.Context, arg CreateStudentParams) error
}

type studentRepository struct{}

func NewStudentRepository() IStudentRepository {
	return &studentRepository{}
}

func (r *studentRepository) CreateStudent(ctx *gin.Context, arg CreateStudentParams) error {
	return global.Db.CreateStudent(ctx, database.CreateStudentParams{
		Code:       arg.Code,
		SubMajorID: arg.SubMajorID,
		UserID:     arg.UserID,
	})
}
