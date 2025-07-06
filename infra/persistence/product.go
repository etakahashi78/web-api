package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"web-api/domain/model"
)

type MySQLProductRepository struct {
	db *sqlx.DB
}

func NewMySQLProductRepository(dsn string) (*MySQLProductRepository, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLProductRepository{db: db}, nil
}

// Save ...
func (repo *MySQLProductRepository) Save(ctx context.Context, product *model.Product) error {
	query := "INSERT INTO products (name, price, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := repo.db.ExecContext(ctx, query, product.Name, product.Price, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = id
	return nil
}

// FindAll ...
func (repo *MySQLProductRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	query := "SELECT id, name, price, created_at, updated_at FROM products"
	var products []*model.Product
	err := repo.db.SelectContext(ctx, &products, query)
	if err != nil {
		msg := "Querying to products table failed."
		slog.Error(msg, "error", err)
		return nil, fmt.Errorf("%s err:%w", msg, err)
	}
	return products, nil
}

// FindByID ...
func (repo *MySQLProductRepository) FindByID(ctx context.Context, id int64) (*model.Product, error) {
	query := "SELECT id, name, price, created_at, updated_at FROM products WHERE id = ?"
	var product model.Product
	err := repo.db.GetContext(ctx, &product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("product was not found.", "id", id)
			return nil, nil
		}
		return nil, fmt.Errorf("scanning failed in querying products table err:%w", err)
	}
	return &product, nil
}
