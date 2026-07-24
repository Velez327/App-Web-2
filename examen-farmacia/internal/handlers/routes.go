// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NuevoRouter registra todas las rutas de la API. Este archivo es el
// contrato HTTP del examen: los tests httptest de acceptance/ atacan
// exactamente estas rutas.
func NuevoRouter(
	medicamentos *MedicamentoHandler,
	clientes *ClienteHandler,
	ventas *VentaHandler,
) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/medicamentos", func(r chi.Router) {
			r.Get("/", medicamentos.Listar)
			r.Post("/", medicamentos.Crear)
		})

		r.Route("/clientes", func(r chi.Router) {
			r.Get("/", clientes.Listar)
			r.Post("/", clientes.Crear)
			r.Get("/{id}", clientes.ObtenerPorID)
		})

		r.Route("/ventas", func(r chi.Router) {
			r.Get("/", ventas.Listar)
			r.Post("/", ventas.Crear)
			r.Get("/{id}", ventas.ObtenerPorID)
			r.Post("/{id}/anular", ventas.Anular)
		})
	})

	return r
}
