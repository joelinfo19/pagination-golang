// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package tutorial

import (
	"database/sql"
)

type Person struct {
	ID        int32
	FirstName sql.NullString
	LastName  sql.NullString
}