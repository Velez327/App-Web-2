package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/services"
	"github.com/joancema/examen-farmacia/internal/storage"
)

func TestVentaService_StockInsuficiente(t *testing.T) {
	// Fakes en memoria — no tocan la base de datos real
	hm := storage.NuevoMedicamentoMemoria()
	cm := storage.NuevoClienteMemoria()
	vm := storage.NuevaVentaMemoria()

	// Crear datos de prueba
	med := models.Medicamento{Nombre: "Paracetamol", PrecioUnitario: 8.5, Stock: 3, Activo: true}
	require.NoError(t, hm.Crear(&med))
	cli := models.Cliente{Nombre: "Ana", Cedula: "1310000001", Telefono: "099"}
	require.NoError(t, cm.Crear(&cli))

	svc := services.NuevaVentaService(vm, hm, cm)

	// Pedir más de lo que hay en stock
	venta := models.Venta{MedicamentoID: med.ID, ClienteID: cli.ID, Cantidad: 10}
	err := svc.Crear(&venta)

	require.ErrorIs(t, err, services.ErrStockInsuficiente,
		"pedir 10 unidades con stock 3 debe devolver ErrStockInsuficiente")
}