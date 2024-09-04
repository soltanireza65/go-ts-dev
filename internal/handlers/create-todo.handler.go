package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/soltanireza65/go-ts-dev/internal/store"
)

type CreateTodoHandler struct {
	todos *[]store.Todo
}

type CreateTodoHandlerParams struct {
	Todos *[]store.Todo
}

func NewCreateTodoHandler(parama CreateTodoHandlerParams) *CreateTodoHandler {
	return &CreateTodoHandler{
		todos: parama.Todos,
	}
}

// func NewCreateTodoHandler(Todos *[]store.Todo) *CreateTodoHandler {
// 	return &CreateTodoHandler{
// 		todos: Todos,
// 	}
// }

func (h *CreateTodoHandler) Excute(w http.ResponseWriter, r *http.Request) {
	todo := store.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	*h.todos = append(*h.todos, todo)

	w.WriteHeader(http.StatusCreated)
}
