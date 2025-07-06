package repository

import (
	"context"

	"web-api/domain/model"
)

type TodoRepository interface {
	Save(context.Context, *model.Todo) error
	Update(context.Context, *model.Todo) (*model.Todo, error)
	FindAll(context.Context) ([]*model.Todo, error)
	FindByID(context.Context, int64) (*model.Todo, error)
}
