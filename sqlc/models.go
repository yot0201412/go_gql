// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc

import (
	"time"
)

type User struct {
	ID        int32
	Name      string
	CreatedAt time.Time
}