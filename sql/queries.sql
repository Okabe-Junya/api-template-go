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

-- name: CreateItem :exec
INSERT INTO items (name, price, description, stock) VALUES (?, ?, ?, ?);

-- name: GetItem :one
SELECT id, name, price, description, stock FROM items WHERE id = ? LIMIT 1;

-- name: ListItems :many
SELECT id, name, price, description, stock FROM items;

-- name: UpdateItem :exec
UPDATE items SET name = ?, price = ?, description = ?, stock = ? WHERE id = ?;

-- name: DeleteItem :exec
DELETE FROM items WHERE id = ?;

-- name: CreateUserItem :exec
INSERT INTO user_items (user_id, item_id, quantity, purchase_date) VALUES (?, ?, ?, ?);

-- name: GetUserItem :one
SELECT user_id, item_id, quantity, purchase_date FROM user_items WHERE user_id = ? AND item_id = ? LIMIT 1;

-- name: ListUserItems :many
SELECT user_id, item_id, quantity, purchase_date FROM user_items;

-- name: UpdateUserItem :exec
UPDATE user_items SET quantity = ?, purchase_date = ? WHERE user_id = ? AND item_id = ?;

-- name: DeleteUserItem :exec
DELETE FROM user_items WHERE user_id = ? AND item_id = ?;
