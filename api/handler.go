package api

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, s UserService) {
	r.Get("/", VersionHandler)
	r.Get("/ping", PongHandler)

	r.Route("/user", func(r chi.Router) {
		r.Post("/list", WrapHandler[ListUserRequest](ListUsersHandler(s)))
		r.Post("/get", WrapHandler[GetUserRequest](GetUserHandler(s)))
		r.Post("/create", WrapHandler[CreateUserRequest](CreateUserHandler(s)))
		r.Post("/update", WrapHandler[UpdateUserRequest](UpdateUserHandler(s)))
		r.Post("/delete", WrapHandler[DeleteUserRequest](DeleteUserHandler(s)))
	})
}
