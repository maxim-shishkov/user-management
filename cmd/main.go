package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"user-management/api"
	"user-management/config"
	"user-management/domain/users/repository"
	"user-management/domain/users/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal("failed to initialize user repository", zap.Error(err))
	}
	defer logger.Sync()

	cfg, err := config.LoadConfig("/config/config.json")
	if err != nil {
		logger.Fatal("failed to load config: %v", zap.Error(err))
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userRepo, err := repository.NewUserRepository(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("failed to initialize user repository", zap.Error(err))
	}
	defer userRepo.Close()

	userService := service.NewUserService(userRepo, logger)
	api.RegisterRoutes(r, userService, logger)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	logger.Info("starting server", zap.String("address", cfg.ServerAddress))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", zap.Error(err))
	}

	logger.Info("server stopped.")
}
