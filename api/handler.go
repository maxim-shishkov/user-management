package api

import (
	"github.com/go-chi/chi/v5"
	"user-management/domain/users"
	"user-management/domain/users/handlers"
)

func RegisterRoutes(r chi.Router, s handlers.UserService) {
	r.Get("/", VersionHandler)
	r.Get("/ping", PongHandler)

	r.Route("/user", func(r chi.Router) {
		r.Post("/list", WrapHandler[users.ListRequest, users.ListResponse](handlers.ListHandler(s)))
		r.Post("/get", WrapHandler[users.GetRequest, users.GetResponse](handlers.GetHandler(s)))
		r.Post("/create", WrapHandler[users.CreateRequest, users.CreateResponse](handlers.CreateHandler(s)))
		r.Post("/update", WrapHandler[users.UpdateRequest, users.UpdateResponse](handlers.UpdateHandler(s)))
		r.Post("/delete", WrapHandler[users.DeleteRequest, users.DeleteResponse](handlers.DeleteHandler(s)))
	})
}
