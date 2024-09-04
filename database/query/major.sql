-- name: GetListMajor :many
SELECT * FROM majors
LIMIT $1 OFFSET $2;

-- name: CountAllMajor :one
SELECT COUNT(*) AS total FROM majors;