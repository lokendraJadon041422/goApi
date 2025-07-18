package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lokendraJadon041422/studentsApi/internal/config"
)

func main() {
	// create config
	config := config.MustLoad()
	// create database

	// create router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	// create serever
	server := &http.Server{
		Addr:    ":" + config.HttpServer.Port,
		Handler: router,
	}
	slog.Info("Server started on :" + config.HttpServer.Port)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", "error", err.Error())
		}
	}()
	<-done
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", "error", err.Error())
	}
	slog.Info("Server shutdown successfully")
}
