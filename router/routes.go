package router

import (
	"github.com/go-chi/chi/v5"

	"web-api/adapter/controllers"
)

// SetupRoutes は HTTP ルーティングを設定します。
func SetupRoutes(
	productController *controllers.ProductController,
	todoController *controllers.TodoController,
) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productController.RegisterProductHandler)
		r.Get("/", productController.ListProductsHandler)
		r.Get("/{id:[0-9]+}", productController.GetProductHandler)
	})
	r.Route("/todos", func(r chi.Router) {
		r.Post("/", todoController.RegisterTodoHandler)
		r.Get("/", todoController.ListTodosHandler)
		r.Get("/{id:[0-9]+}", todoController.GetTodoHandler)
	})

	return r
}
