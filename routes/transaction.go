package routes

import (
	"dbf_api/services"

	"github.com/go-chi/chi/v5"

)

func RegisterHTTPEndpoints(r chi.Router, service *services.Service) chi.Router {
	r.Get("/", service.ListTransactions)
	r.Post("/", service.CreateRecord)
	r.Get("/{id}", service.GetById)
	r.Put("/{id}", service.UpdateRecord)
	r.Delete("/{id}", service.DeleteRecord)

    return r
}
