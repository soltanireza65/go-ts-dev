package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/soltanireza65/go-ts-dev/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	serverCtx, serverStopCancelFunc := context.WithCancel(context.Background())

	killSigChan := make(chan os.Signal, 1)
	signal.Notify(killSigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		sig := <-killSigChan
		logger.Info("got kill signal - shutting down", slog.String("signal", sig.String()))

		shutdownCtx, shutdownCancelFunc := context.WithTimeout(serverCtx, 5*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				// print("graceful shutdown timed out - forcing exit\n")
				// os.Exit(1)
				log.Fatal("shutdown deadline exceeded")
			}
		}()

		err := srv.Shutdown(shutdownCtx)

		if err != nil {
			log.Fatal(err)
		}

		serverStopCancelFunc()
		logger.Info("server shutting down")
		shutdownCancelFunc()
	}()

	go func() {
		err := srv.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	r.Get("/healthcheck", handlers.NewHealthcheckHandler().Excute)

	logger.Info("read to work")

	<-serverCtx.Done()
}
