// Package main as main entrypoint of application
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrjxtr/psalms/internal/config"
	"github.com/mrjxtr/psalms/internal/routes"
)

func main() {
	slog.Info("Starting server")

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	r, err := routes.SetupRouter()
	if err != nil {
		slog.Error("Error setting up router", "error", err)
		os.Exit(1)
	}
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		slog.Info("Server started", "port", cfg.Port)
		slog.Info(
			"Check server health",
			"url",
			fmt.Sprintf("http://127.0.0.1:%s/ping", cfg.Port),
		)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	slog.Info("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Graceful shutdown failed...", "error", err)
		os.Exit(1)
	} else {
		slog.Info("server gracefully stopped")
	}
}
