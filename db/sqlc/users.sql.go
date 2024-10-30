package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  alias, email, password, first_name, last_name, cpf, phone_number
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING user_id, alias, email, password, create_time, first_name, last_name, cpf, phone_number
`

type CreateUserParams struct {
	Alias       string      `json:"alias"`
	Email       interface{} `json:"email"`
	Password    string      `json:"password"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Cpf         interface{} `json:"cpf"`
	PhoneNumber interface{} `json:"phone_number"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Alias,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.Cpf,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Alias,
		&i.Email,
		&i.Password,
		&i.CreateTime,
		&i.FirstName,
		&i.LastName,
		&i.Cpf,
		&i.PhoneNumber,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_ID = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, alias, email, password, create_time, first_name, last_name, cpf, phone_number FROM users
WHERE user_ID = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Alias,
		&i.Email,
		&i.Password,
		&i.CreateTime,
		&i.FirstName,
		&i.LastName,
		&i.Cpf,
		&i.PhoneNumber,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, alias, email, password, create_time, first_name, last_name, cpf, phone_number FROM users
ORDER BY user_ID
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Alias,
			&i.Email,
			&i.Password,
			&i.CreateTime,
			&i.FirstName,
			&i.LastName,
			&i.Cpf,
			&i.PhoneNumber,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
  set first_name = $2
WHERE user_ID = $1
RETURNING user_id, alias, email, password, create_time, first_name, last_name, cpf, phone_number
`

type UpdateUserParams struct {
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.UserID, arg.FirstName)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Alias,
		&i.Email,
		&i.Password,
		&i.CreateTime,
		&i.FirstName,
		&i.LastName,
		&i.Cpf,
		&i.PhoneNumber,
	)
	return i, err
}
