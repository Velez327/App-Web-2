package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/services"
)
// TAREA (CP3): Implemente VentaHandler.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Mapeo de errores de dominio a status codes (los tests lo verifican):
//       ErrDatosInvalidos     -> 422 Unprocessable Entity
//       ErrReferenciaInvalida -> 422 Unprocessable Entity
//       ErrStockInsuficiente  -> 409 Conflict
//       ErrEstadoInvalido     -> 409 Conflict
//       ErrNoEncontrado       -> 404 Not Found
//       cualquier otro error  -> 500 Internal Server Error
type VentaHandler struct {
	servicio *services.VentaService
}

func NuevaVentaHandler(s *services.VentaService) *VentaHandler {
	return &VentaHandler{servicio: s}
}

func (h *VentaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 201 con la venta creado (incluido el total).
	var venta models.Venta
	if err := json.NewDecoder(r.Body).Decode(&venta); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.servicio.Crear(&venta); err != nil {
		switch {
		case errors.Is(err, services.ErrReferenciaInvalida):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		case errors.Is(err, services.ErrStockInsuficiente):
			RespondError(w, http.StatusConflict, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondJSON(w, http.StatusCreated, venta)
}

func (h *VentaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200 con la lista.
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *VentaHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200; no existe -> 404.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	venta, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, services.ErrNoEncontrado):
			RespondError(w, http.StatusNotFound, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondJSON(w, http.StatusOK, venta)
}

func (h *VentaHandler) Anular(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200; estado inválido -> 409; no existe -> 404.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.servicio.Anular(uint(id)); err != nil {
		switch {
		case errors.Is(err, services.ErrNoEncontrado):
			RespondError(w, http.StatusNotFound, err.Error())
		case errors.Is(err, services.ErrEstadoInvalido):
			RespondError(w, http.StatusConflict, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondJSON(w, http.StatusOK, nil)
}
