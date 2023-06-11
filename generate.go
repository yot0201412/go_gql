//go:generate sqlc-dev generate --file sqlc.yaml
//go:generate go run github.com/99designs/gqlgen generate
package main

import (
	_ "github.com/99designs/gqlgen"
)
