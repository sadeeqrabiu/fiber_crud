// Code generated by sqlc. DO NOT EDIT.
// source: todos.sql

package db

import (
	"context"
)

const getAllTodos = `-- name: GetAllTodos :many
SELECT id, name, completed FROM todos
`

func (q *Queries) GetAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.query(ctx, q.getAllTodosStmt, getAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Name, &i.Completed); err != nil {
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
