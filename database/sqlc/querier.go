// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
)

type Querier interface {
	GetRole(ctx context.Context, id int64) (Role, error)
}

var _ Querier = (*Queries)(nil)
