package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"user-management/api"
	"user-management/config"
	"user-management/domain/users/repository"
	"user-management/domain/users/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.LoadConfig("/config/config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userRepo, err := repository.NewUserRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to initialize user repository: %v", err)
	}
	defer userRepo.Close()

	userService := service.NewUserService(userRepo)
	api.RegisterRoutes(r, userService)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server stopped.")
}
