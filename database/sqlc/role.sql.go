// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: role.sql

package database

import (
	"context"
)

const getRole = `-- name: GetRole :one
SELECT id, name FROM roles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, id int64) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRole, id)
	var i Role
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
