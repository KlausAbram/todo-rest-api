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
	Delete(userId, listId int) error
	Update(userId, listId int, input todoapi.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, input todoapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoapi.TodoItem, error)
	GetById(userId, itemId int) (todoapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoapi.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
