-- name: GetUsers :many
select *
from users;

-- name: GetUser :one
select *
from users
where id = $1;

-- name: CreateUser :exec
insert into users (name) values ($1);