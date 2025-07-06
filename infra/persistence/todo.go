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

type MySQLTodoRepository struct {
	db *sqlx.DB
}

func NewMySQLTodoRepository(dsn string) (*MySQLTodoRepository, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open failed. err:%w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping failed. err:%w", err)
	}
	return &MySQLTodoRepository{db: db}, nil
}

func (repo *MySQLTodoRepository) Save(ctx context.Context, todo *model.Todo) error {
	query := `INSERT INTO todo (name, content, done, created_at, updated_at) VALUES(?, ?, ?, ?, ?)`
	result, err := repo.db.ExecContext(ctx, query, todo.Name, todo.Content, todo.Done, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save. err:%w", err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected failed. err:%w", err)
	}

	return nil
}

// Update ...
func (repo *MySQLTodoRepository) Update(ctx context.Context, t *model.Todo) (*model.Todo, error) {
	//query := "UPDATE `todos` SET name, content, done, created_at, updated_at WHERE id = ?"
	//
	//result, err := repo.db.ExecContext(ctx, query, t.Name, t.Content, t.Done, t.CreatedAt, tß.UpdatedAt)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to save. err:%w", err)
	//}
	//_, err = result.RowsAffected()
	//if err != nil {
	//	return nil, fmt.Errorf("RowsAffected failed. err:%w", err)
	//}

	return nil, nil
}

// FindAll ...
func (repo *MySQLTodoRepository) FindAll(ctx context.Context) (todos []*model.Todo, err error) {
	query := `SELECT id, name, content, done, created_at, updated_at FROM todos`

	err = repo.db.SelectContext(ctx, &todos, query)
	if err != nil {
		err = fmt.Errorf("SelectContext failed. err:%w", err)
		return
	}

	return
}

// FindByID ...
func (repo *MySQLTodoRepository) FindByID(ctx context.Context, id int64) (todo *model.Todo, err error) {
	query := `SELECT id, name, content, done, created_at, updated_at FROM todos WHERE id = ?`

	err = repo.db.GetContext(ctx, &todo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("todo was not found.", "id", id)
			return
		}
		err = fmt.Errorf("GetContext failed. err:%w", err)
		return
	}

	return
}
