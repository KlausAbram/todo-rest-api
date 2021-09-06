package service

import (
	todoapi "github.com/klaus-abram/todo-rest-api"
	"github.com/klaus-abram/todo-rest-api/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, input todoapi.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, input)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todoapi.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todoapi.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input todoapi.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, input)
}
