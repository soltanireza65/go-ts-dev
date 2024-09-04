package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/soltanireza65/go-ts-dev/internal/handlers"
	"github.com/soltanireza65/go-ts-dev/internal/store"
)

func main() {

	todos := []store.Todo{}

	r := chi.NewRouter()

	// Use chi's logger and recover middlewares for better error handling
	r.Use(middleware.Logger)

	r.Get("/healthcheck", handlers.NewHealthcheckHandler().Excute)

	// r.Post("/todos", handlers.NewCreateTodoHandler(&todos).Excute)

	r.Post("/todos", handlers.NewCreateTodoHandler(handlers.CreateTodoHandlerParams{Todos: &todos}).Excute)

	r.Get("/todos", handlers.NewListTodosHandler(handlers.ListTodosHandlerParams{Todos: &todos}).Excute)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Create a channel to listen for OS signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Server is running on :8080")

		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	fmt.Println("Press Ctrl+C to stop the server")

	// Wait for signals to gracefully shut down the server
	<-sigCh

	fmt.Println("Shutting down the server...")

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}
