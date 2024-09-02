-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (name, user_type, password, email, phone_number)
VALUES ($1, $2, $3, $4, $5);

-- name: CreateUserAndReturnId :one
INSERT INTO users (name, user_type, password, email, phone_number)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;