// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Todo struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Completed sql.NullBool `json:"completed"`
}
