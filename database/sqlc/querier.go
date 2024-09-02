// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
)

type Querier interface {
	CreateStudent(ctx context.Context, arg CreateStudentParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	CreateUserAndReturnId(ctx context.Context, arg CreateUserAndReturnIdParams) (int64, error)
	GetRole(ctx context.Context, id int64) (Role, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
}

var _ Querier = (*Queries)(nil)
