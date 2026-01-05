package repository

import (
	"database/sql"
	"gachibank/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
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

func NewRepository(db *sql.DB) *Repository { // Конструктор для запуска в main
	return &Repository{
		Authorization: NewAuthSQL(db),
	}
}
