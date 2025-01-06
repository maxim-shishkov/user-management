package main

import (
	"log"
	"net/http"

	"user-management/api"
	"user-management/config"
	"user-management/repository"
	"user-management/service"

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

	userRepo, err := repository.NewUserRepository(cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	userService := service.NewUserService(userRepo)
	api.RegisterRoutes(r, userService)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
