package main

import (
	"database/sql"
	"fmt"
	"log"

	"user-management/api"
	"user-management/config"
	"user-management/repository"
	"user-management/server"
	"user-management/service"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Database connection
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize layers
	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	handler := api.NewHandler(srv)
	s := server.NewServer(handler)

	// Start server
	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := s.Start(cfg.Server.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
