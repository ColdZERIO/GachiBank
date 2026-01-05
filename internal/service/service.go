package service

import (
	"github.com/TifaLuv/GolangServer/domain"
	"github.com/TifaLuv/GolangServer/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(rep *repository.Repository) *Service { // Конструктор для запуска в main
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
