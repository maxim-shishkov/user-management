package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"user-management/service"
)

func RegisterRoutes(r chi.Router, s *service.UserService) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/get", WrapHandler[service.GetUserRequest](s.GetUser))
		r.Post("/create", WrapHandler[service.CreateUserRequest](s.CreateUser))
		r.Post("/update", WrapHandler[service.UpdateUserRequest](s.UpdateUser))
		r.Post("/delete", WrapHandler[service.DeleteUserRequest](s.DeleteUser))
	})
}
