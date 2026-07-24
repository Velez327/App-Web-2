package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"github.com/joancema/examen-farmacia/internal/handlers"
	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/services"
	"github.com/joancema/examen-farmacia/internal/storage"
)

func TestVentaHandler_Crear_201(t *testing.T) {
	hm := storage.NuevoMedicamentoMemoria()
	cm := storage.NuevoClienteMemoria()
	vm := storage.NuevaVentaMemoria()

	med := models.Medicamento{Nombre: "Paracetamol", PrecioUnitario: 8.5, Stock: 10, Activo: true}
	require.NoError(t, hm.Crear(&med))
	cli := models.Cliente{Nombre: "Ana", Cedula: "1310000001", Telefono: "099"}
	require.NoError(t, cm.Crear(&cli))

	svc := services.NuevaVentaService(vm, hm, cm)
	h := handlers.NuevaVentaHandler(svc)

	r := chi.NewRouter()
	r.Post("/api/v1/ventas", h.Crear)

	body := fmt.Sprintf(`{"medicamento_id":%d,"cliente_id":%d,"cantidad":2}`, med.ID, cli.ID)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ventas", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code,
		"POST /ventas válido debe responder 201. Body: %s", rec.Body.String())
}