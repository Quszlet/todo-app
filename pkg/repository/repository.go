package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(bd *sqlx.DB) *Repository {
	return &Repository{}
}
