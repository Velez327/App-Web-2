// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/services"
)

// MedicamentoHandler expone la Entidad A por HTTP.
// Está completo: observe cómo decodifica el body, llama al service y
// MAPEA los errores de dominio a status codes. Ese mapeo es exactamente
// lo que usted debe replicar en sus propios handlers.
type MedicamentoHandler struct {
	servicio *services.MedicamentoService
}

func NuevoMedicamentoHandler(s *services.MedicamentoService) *MedicamentoHandler {
	return &MedicamentoHandler{servicio: s}
}

func (h *MedicamentoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *MedicamentoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var medicamento models.Medicamento
	if err := json.NewDecoder(r.Body).Decode(&medicamento); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if err := h.servicio.Crear(&medicamento); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, medicamento)
}
