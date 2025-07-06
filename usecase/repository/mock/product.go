package mock

import (
	"context"

	"web-api/domain/model/entity"
)

type ProductMock struct {
	SaveFunc     func(context.Context) error
	FindAllFunc  func(context.Context) ([]*entity.Product, error)
	FindByIDFunc func(context.Context, int64) (*entity.Product, error)
}

func (m *ProductMock) Save(ctx context.Context) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx)
	}
	return nil
}

func (m *ProductMock) FindAll(ctx context.Context) ([]*entity.Product, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc(ctx)
	}
	return nil, nil
}

func (m *ProductMock) FindByID(ctx context.Context, productID int64) (*entity.Product, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(ctx, productID)
	}
	return nil, nil
}
