package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/soltanireza65/go-ts-dev/internal/store"
)

type ListTodosHandler struct {
	todos *[]store.Todo
}

type ListTodosHandlerParams struct {
	Todos *[]store.Todo
}

func NewListTodosHandler(parama ListTodosHandlerParams) *ListTodosHandler {
	return &ListTodosHandler{
		todos: parama.Todos,
	}
}

func (h *ListTodosHandler) Excute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.todos)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
