package repository

import (
	database "github.com/api/database/sqlc"
	"github.com/api/global"
	"github.com/gin-gonic/gin"
)

type GetListMajorParams database.GetListMajorParams

type IMajorRepository interface {
	GetListMajor(ctx *gin.Context, arg GetListMajorParams) ([]database.Major, error)
	CountAllMajor(ctx *gin.Context) (int64, error)
}

type majorRepository struct{}

func NewMajorRepository() IMajorRepository {
	return &majorRepository{}
}

func (r *majorRepository) GetListMajor(ctx *gin.Context, arg GetListMajorParams) ([]database.Major, error) {
	return global.Db.GetListMajor(ctx, database.GetListMajorParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
	})
}

func (r *majorRepository) CountAllMajor(ctx *gin.Context) (int64, error) {
	return global.Db.CountAllMajor(ctx)
}