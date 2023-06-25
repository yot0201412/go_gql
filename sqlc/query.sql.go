// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :exec
insert into todos (user_id, text) values ($1, $2)
`

type CreateTodoParams struct {
	UserID int32
	Text   string
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) error {
	_, err := q.db.ExecContext(ctx, createTodo, arg.UserID, arg.Text)
	return err
}

const createUser = `-- name: CreateUser :exec
insert into users (name) values ($1)
`

func (q *Queries) CreateUser(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createUser, name)
	return err
}

const getTodosByUserID = `-- name: GetTodosByUserID :many
select id, user_id, text, done, created_at
from todos
where
    true
    and (
        user_id = $1 OR $1 IS NULL
    -- ここを逆で書いてしまうと、動かないので注意
    )
`

func (q *Queries) GetTodosByUserID(ctx context.Context, userID sql.NullInt32) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getTodosByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Text,
			&i.Done,
			&i.CreatedAt,
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

const getUser = `-- name: GetUser :one
select id, name, created_at
from users
where id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select id, name, created_at
from users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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

const selectNameFromJson = `-- name: SelectNameFromJson :one
select
    id,
    json_data->>'name' as name_json,
    'name' || ' first' as name_first
from json_table
where id = $1
`

type SelectNameFromJsonRow struct {
	ID        int32
	NameJson  sql.NullString
	NameFirst sql.NullString
}

func (q *Queries) SelectNameFromJson(ctx context.Context, id int32) (SelectNameFromJsonRow, error) {
	row := q.db.QueryRowContext(ctx, selectNameFromJson, id)
	var i SelectNameFromJsonRow
	err := row.Scan(&i.ID, &i.NameJson, &i.NameFirst)
	return i, err
}

const selectTJson = `-- name: SelectTJson :one
select id, json_data
from json_table
where id = $1
`

func (q *Queries) SelectTJson(ctx context.Context, id int32) (JsonTable, error) {
	row := q.db.QueryRowContext(ctx, selectTJson, id)
	var i JsonTable
	err := row.Scan(&i.ID, &i.JsonData)
	return i, err
}
