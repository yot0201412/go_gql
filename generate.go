// -- go:generate pg_dump -U postgres -p 5432 -d go_gql -h localhost -T 'gorp*' -s > schema.sql
//
//go:generate sqlc generate --file sqlc.yaml
//go:generate go run github.com/99designs/gqlgen generate
package main

import (
	_ "github.com/99designs/gqlgen"
)
