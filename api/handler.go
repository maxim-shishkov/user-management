package api

import (
	"github.com/go-chi/chi/v5"
	"user-management/domain/users/handlers"
)

func RegisterRoutes(r chi.Router, s handlers.UserService) {
	r.Get("/", VersionHandler)
	r.Get("/ping", PongHandler)

	r.Route("/user", func(r chi.Router) {
		r.Post("/list", WrapHandler[handlers.ListRequest, handlers.ListResponse](handlers.ListHandler(s)))
		r.Post("/get", WrapHandler[handlers.GetRequest, handlers.GetResponse](handlers.GetHandler(s)))
		r.Post("/create", WrapHandler[handlers.CreateRequest, handlers.CreateResponse](handlers.CreateHandler(s)))
		r.Post("/update", WrapHandler[handlers.UpdateRequest, handlers.UpdateResponse](handlers.UpdateHandler(s)))
		r.Post("/delete", WrapHandler[handlers.DeleteRequest, handlers.DeleteResponse](handlers.DeleteHandler(s)))
	})
}
