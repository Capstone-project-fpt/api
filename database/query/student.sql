-- name: CreateStudent :exec
INSERT INTO students (code, sub_major_id, user_id)
VALUES ($1, $2, $3);