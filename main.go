package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	ctrl "web-api/adapter/controllers"
	"web-api/config"
	"web-api/infra"
	"web-api/router"
	"web-api/usecase/interactor"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		log.Fatal(err)
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config.Config) error {
	repos := infra.NewRepositories(cfg)
	// DI
	pi := interactor.NewProductInteractor(repos.ProductRepository)
	productController := ctrl.NewProductController(pi)
	ti := interactor.NewTodoInteractor(repos.TodoRepository)
	todoController := ctrl.NewTodoController(ti)

	routes := router.SetupRoutes(
		productController,
		todoController,
	)

	slog.Info("Server started on port 8080")
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		slog.Error("failed to start server:", "err", err)
	}

	return err
}
