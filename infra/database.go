package infra

import (
	"database/sql"
	"log"
	"log/slog"

	"web-api/config"
	"web-api/infra/persistence"
	"web-api/usecase/repository"
)

type Repositories struct {
	primaryDB         *sql.DB
	replicaDB         *sql.DB
	ProductRepository repository.ProductRepository
	TodoRepository    repository.TodoRepository
}

func NewRepositories(cfg *config.Config) *Repositories {
	primaryDB, err := sql.Open("mysql", cfg.DbDsnPrimary)
	if err != nil {
		log.Fatalf("failed to connect to primary MySQL: %v", err)
	}
	if err := primaryDB.Ping(); err != nil {
		log.Fatalf("failed to ping primary MySQL: %v", err)
	}

	replicaDB, err := sql.Open("mysql", cfg.DbDsnReplica)
	if err != nil {
		slog.Info("failed to connect to replica MySQL, using primary instead", "err", err)
		replicaDB = primaryDB
	} else {
		if err := replicaDB.Ping(); err != nil {
			slog.Info("failed to ping replica MySQL, using primary instead", "err", err)
			replicaDB = primaryDB
		}
	}

	// 商品テーブル
	productRepo, err := persistence.NewMySQLProductRepository(cfg.DbDsnPrimary)
	if err != nil {
		log.Fatalf("failed to connect to MySQL for ProductRepository: %v", err)
	}
	// TODOテーブル
	todoRepo, err := persistence.NewMySQLTodoRepository(cfg.DbDsnPrimary)
	if err != nil {
		log.Fatalf("failed to connect to MySQL for TodoRepository: %v", err)
	}

	return &Repositories{
		primaryDB:         primaryDB,
		replicaDB:         replicaDB,
		ProductRepository: productRepo,
		TodoRepository:    todoRepo,
	}
}

func (r *Repositories) Close() {
	if r.primaryDB != nil {
		err := r.primaryDB.Close()
		if err != nil {
			slog.Error("failed to close primaryDB connection.", "err", err)
			return
		}
	}
	if r.replicaDB != nil && r.replicaDB != r.primaryDB {
		err := r.replicaDB.Close()
		if err != nil {
			slog.Error("failed to close replicaDB connection.", "err", err)
			return
		}
	}
}
