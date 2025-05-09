-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: DeleteUsers :exec
TRUNCATE users RESTART IDENTITY CASCADE;