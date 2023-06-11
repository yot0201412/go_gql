-- name: GetUsers :many
select *
from users;

-- name: GetUser :one
select *
from users
where id = $1;

-- name: CreateUser :exec
insert into users (name) values ($1);

-- name: CreateTodo :exec
insert into todos (user_id, text) values ($1, $2);

-- name: GetTodosByUserID :many
select *
from todos
where
    true
    and (
        user_id = sqlc.narg('user_id') OR sqlc.narg('user_id') IS NULL
    -- ここを逆で書いてしまうと、動かないので注意
    )
;

-- name: SelectTJson :one
select *
from json_table
where id = $1;