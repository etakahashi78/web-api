package interactor

import (
	"context"

	"web-api/usecase/repository"
)

type TodoInteractor any

type todoInteractor struct {
	todoRepo repository.TodoRepository
}

func NewTodoInteractor(repo repository.TodoRepository) TodoInteractor {
	return todoInteractor{todoRepo: repo}
}

func (i *todoInteractor) RegisterTodo(ctx context.Context, name, content string, done bool) {

}
