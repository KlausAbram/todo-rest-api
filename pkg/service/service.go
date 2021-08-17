package service

import (
	todoapi "github.com/klaus-abram/todo-rest-api"
	"github.com/klaus-abram/todo-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todoapi.TodoList) (int, error)
	GetAll(userId int) ([]todoapi.TodoList, error)
	GetById(userId, listId int) (todoapi.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
