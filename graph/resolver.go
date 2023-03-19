package graph

import (
	sqlc "github.com/yot0201412/go_gql/sqlc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repo *sqlc.Queries
}
