package api

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, s UserService) {
	r.Get("/", VersionHandler)
	r.Get("/ping", PongHandler)

	r.Route("/user", func(r chi.Router) {
		r.Post("/list", WrapHandler[ListRequest, ListResponse](ListHandler(s)))
		r.Post("/get", WrapHandler[GetRequest, GetResponse](GetHandler(s)))
		r.Post("/create", WrapHandler[CreateRequest, CreateResponse](CreateHandler(s)))
		r.Post("/update", WrapHandler[UpdateRequest, UpdateResponse](UpdateHandler(s)))
		r.Post("/delete", WrapHandler[DeleteRequest, DeleteResponse](DeleteHandler(s)))
	})
}
