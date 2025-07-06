package controllers

import (
	"net/http"

	"web-api/domain/model"
	"web-api/usecase/interactor"
)

type TodoController struct {
	todoInteractor interactor.TodoInteractor
}

func NewTodoController(i interactor.TodoInteractor) *TodoController {
	return &TodoController{todoInteractor: i}
}

// RegisterTodoHandler ...
func (c *TodoController) RegisterTodoHandler(w http.ResponseWriter, r *http.Request) {

	todo := &model.Todo{}

	ResponseJSON(w, http.StatusCreated, todo)
}

// ListTodosHandler ...
func (c *TodoController) ListTodosHandler(w http.ResponseWriter, r *http.Request) {

	todo := &model.Todo{}

	ResponseJSON(w, http.StatusOK, todo)
}

// GetTodoHandler ...
func (c *TodoController) GetTodoHandler(w http.ResponseWriter, r *http.Request) {

	todo := &model.Todo{}

	ResponseJSON(w, http.StatusOK, todo)
}
