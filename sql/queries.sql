-- name: CreateUser :exec
INSERT INTO users (name, email) VALUES (?, ?);

-- name: GetUser :one
SELECT id, name, email FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT id, name, email FROM users;

-- name: UpdateUser :exec
UPDATE users SET name = ?, email = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
