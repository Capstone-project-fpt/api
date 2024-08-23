-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (name, user_type, password, email, code, sub_major_id, capstone_group_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);