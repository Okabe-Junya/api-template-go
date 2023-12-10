// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
)

const createItem = `-- name: CreateItem :exec
INSERT INTO items (name, price, description, stock) VALUES (?, ?, ?, ?)
`

type CreateItemParams struct {
	Name        string         `json:"name"`
	Price       string         `json:"price"`
	Description sql.NullString `json:"description"`
	Stock       int32          `json:"stock"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) error {
	_, err := q.db.ExecContext(ctx, createItem,
		arg.Name,
		arg.Price,
		arg.Description,
		arg.Stock,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name, email) VALUES (?, ?)
`

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Email)
	return err
}

const createUserItem = `-- name: CreateUserItem :exec
INSERT INTO user_items (user_id, item_id, quantity, purchase_date) VALUES (?, ?, ?, ?)
`

type CreateUserItemParams struct {
	UserID       int32        `json:"user_id"`
	ItemID       int32        `json:"item_id"`
	Quantity     int32        `json:"quantity"`
	PurchaseDate sql.NullTime `json:"purchase_date"`
}

func (q *Queries) CreateUserItem(ctx context.Context, arg CreateUserItemParams) error {
	_, err := q.db.ExecContext(ctx, createUserItem,
		arg.UserID,
		arg.ItemID,
		arg.Quantity,
		arg.PurchaseDate,
	)
	return err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items WHERE id = ?
`

func (q *Queries) DeleteItem(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const deleteUserItem = `-- name: DeleteUserItem :exec
DELETE FROM user_items WHERE user_id = ? AND item_id = ?
`

type DeleteUserItemParams struct {
	UserID int32 `json:"user_id"`
	ItemID int32 `json:"item_id"`
}

func (q *Queries) DeleteUserItem(ctx context.Context, arg DeleteUserItemParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserItem, arg.UserID, arg.ItemID)
	return err
}

const getItem = `-- name: GetItem :one
SELECT id, name, price, description, stock FROM items WHERE id = ? LIMIT 1
`

func (q *Queries) GetItem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Stock,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email FROM users WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const getUserItem = `-- name: GetUserItem :one
SELECT user_id, item_id, quantity, purchase_date FROM user_items WHERE user_id = ? AND item_id = ? LIMIT 1
`

type GetUserItemParams struct {
	UserID int32 `json:"user_id"`
	ItemID int32 `json:"item_id"`
}

func (q *Queries) GetUserItem(ctx context.Context, arg GetUserItemParams) (UserItem, error) {
	row := q.db.QueryRowContext(ctx, getUserItem, arg.UserID, arg.ItemID)
	var i UserItem
	err := row.Scan(
		&i.UserID,
		&i.ItemID,
		&i.Quantity,
		&i.PurchaseDate,
	)
	return i, err
}

const listItems = `-- name: ListItems :many
SELECT id, name, price, description, stock FROM items
`

func (q *Queries) ListItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Description,
			&i.Stock,
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

const listUserItems = `-- name: ListUserItems :many
SELECT user_id, item_id, quantity, purchase_date FROM user_items
`

func (q *Queries) ListUserItems(ctx context.Context) ([]UserItem, error) {
	rows, err := q.db.QueryContext(ctx, listUserItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserItem
	for rows.Next() {
		var i UserItem
		if err := rows.Scan(
			&i.UserID,
			&i.ItemID,
			&i.Quantity,
			&i.PurchaseDate,
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

const listUsers = `-- name: ListUsers :many
SELECT id, name, email FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
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

const updateItem = `-- name: UpdateItem :exec
UPDATE items SET name = ?, price = ?, description = ?, stock = ? WHERE id = ?
`

type UpdateItemParams struct {
	Name        string         `json:"name"`
	Price       string         `json:"price"`
	Description sql.NullString `json:"description"`
	Stock       int32          `json:"stock"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) error {
	_, err := q.db.ExecContext(ctx, updateItem,
		arg.Name,
		arg.Price,
		arg.Description,
		arg.Stock,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET name = ?, email = ? WHERE id = ?
`

type UpdateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.Email, arg.ID)
	return err
}

const updateUserItem = `-- name: UpdateUserItem :exec
UPDATE user_items SET quantity = ?, purchase_date = ? WHERE user_id = ? AND item_id = ?
`

type UpdateUserItemParams struct {
	Quantity     int32        `json:"quantity"`
	PurchaseDate sql.NullTime `json:"purchase_date"`
	UserID       int32        `json:"user_id"`
	ItemID       int32        `json:"item_id"`
}

func (q *Queries) UpdateUserItem(ctx context.Context, arg UpdateUserItemParams) error {
	_, err := q.db.ExecContext(ctx, updateUserItem,
		arg.Quantity,
		arg.PurchaseDate,
		arg.UserID,
		arg.ItemID,
	)
	return err
}
