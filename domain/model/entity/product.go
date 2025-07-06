package entity

import (
	"errors"
	"time"
)

type Product struct {
	ID        int64
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if price < 0 {
		return nil, errors.New("product price cannot be negative")
	}
	now := time.Now()
	return &Product{
		Name:      name,
		Price:     price,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (p *Product) UpdatePrice(newPrice float64) error {
	if newPrice < 0 {
		return errors.New("product price cannot be negative")
	}
	p.Price = newPrice
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Product) UpdateName(newName string) error {
	if newName == "" {
		return errors.New("product name cannot be empty")
	}
	p.Name = newName
	p.UpdatedAt = time.Now()
	return nil
}
