package repository

import (
	"context"

	"web-api/domain/model"
)

type ProductRepository interface {
	Save(context.Context, *model.Product) error
	FindAll(context.Context) ([]*model.Product, error)
	FindByID(context.Context, int64) (*model.Product, error)
}
