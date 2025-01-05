package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"user-management/api"
)

type Server struct {
	handler *api.Handler
}

func NewServer(handler *api.Handler) *Server {
	return &Server{handler: handler}
}

func (s *Server) Start(port string) error {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", s.handler.ListUsers)         // GET /users
		r.Get("/{id}", s.handler.GetUserByID)   // GET /users/{id}
		r.Post("/", s.handler.CreateUser)       // POST /users
		r.Put("/{id}", s.handler.UpdateUser)    // PUT /users/{id}
		r.Delete("/{id}", s.handler.DeleteUser) // DELETE /users/{id}
	})

	// Start server
	return http.ListenAndServe(":"+port, r)
}
