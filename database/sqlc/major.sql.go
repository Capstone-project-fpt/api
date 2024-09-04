// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: major.sql

package database

import (
	"context"
)

const countAllMajor = `-- name: CountAllMajor :one
SELECT COUNT(*) AS total FROM majors
`

func (q *Queries) CountAllMajor(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAllMajor)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const getListMajor = `-- name: GetListMajor :many
SELECT id, name, created_at, updated_at FROM majors
LIMIT $1 OFFSET $2
`

type GetListMajorParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetListMajor(ctx context.Context, arg GetListMajorParams) ([]Major, error) {
	rows, err := q.db.QueryContext(ctx, getListMajor, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Major{}
	for rows.Next() {
		var i Major
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
