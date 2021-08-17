package repository

import (
	"github.com/jmoiron/sqlx"
	todoapi "github.com/klaus-abram/todo-rest-api"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemTable   = "todo_items"
	listsItemsTable = "lists_items"
)

type Authorization interface {
	CreateUser(user todoapi.User) (int, error)
	GetUser(username, password string) (todoapi.User, error)
}

type TodoList interface {
	Create(userId int, list todoapi.TodoList) (int, error)
	GetAll(userId int) ([]todoapi.TodoList, error)
	GetById(userId, listId int) (todoapi.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
