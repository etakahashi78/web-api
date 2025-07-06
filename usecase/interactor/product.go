package interactor

import (
	"context"
	"fmt"
	"time"

	"web-api/domain/model"
	"web-api/usecase/repository"
)

// ProductInteractor は商品のユースケースを処理するインターフェースです。
type ProductInteractor interface {
	RegisterProduct(cxt context.Context, name string, price float64) (*model.Product, error)
	ListProducts(context.Context) ([]*model.Product, error)
	GetProduct(context.Context, int64) (*model.Product, error)
}

// productInteractor ProductInteractor の実装
type productInteractor struct {
	productRepo repository.ProductRepository
}

func NewProductInteractor(repo repository.ProductRepository) ProductInteractor {
	return &productInteractor{productRepo: repo}
}

// RegisterProduct ...
func (pi *productInteractor) RegisterProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	now := time.Now()
	newProduct := &model.Product{
		Name:      name,
		Price:     price,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := pi.productRepo.Save(ctx, newProduct)
	if err != nil {
		return nil, fmt.Errorf("productRepo.Save failed. err:%w", err)
	}
	return newProduct, nil
}

// ListProducts ...
func (pi *productInteractor) ListProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := pi.productRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("productRepo.FindAll failed. err:%w", err)
	}
	return products, nil
}

// GetProduct ...
func (pi *productInteractor) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	product, err := pi.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("productRepo.FindByID failed. err:%w", err)
	}
	return product, nil
}
