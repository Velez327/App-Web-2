// ARCHIVO BLOQUEADO — NO MODIFICAR
//
// Las 5 reglas de negocio se verifican aquí usando los repositorios EN MEMORIA
// (ya implementados en el repo base) como fakes. Así, estos tests miden solo
// la lógica de su VentaService, sin depender de su implementación GORM.
package acceptance

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/services"
	"github.com/joancema/examen-farmacia/internal/storage"
)

type entornoReglas struct {
	svc          *services.VentaService
	medicamentos *storage.MedicamentoMemoria
	clientes     *storage.ClienteMemoria
	ventas   *storage.VentaMemoria
	principal      models.Medicamento
	ana          models.Cliente
}

func nuevoEntornoReglas(t *testing.T) entornoReglas {
	t.Helper()
	hm := storage.NuevoMedicamentoMemoria()
	cm := storage.NuevoClienteMemoria()
	am := storage.NuevaVentaMemoria()

	principal := models.Medicamento{Nombre: "Paracetamol 500mg", PrecioUnitario: 8.5, Stock: 10, Activo: true}
	require.NoError(t, hm.Crear(&principal))
	ana := models.Cliente{Nombre: "Ana Zambrano", Cedula: "1310000001", Telefono: "0990000001"}
	require.NoError(t, cm.Crear(&ana))

	return entornoReglas{
		svc:          services.NuevaVentaService(am, hm, cm),
		medicamentos: hm,
		clientes:     cm,
		ventas:   am,
		principal:      principal,
		ana:          ana,
	}
}

// R1: no se crea una venta si el medicamento no existe o está inactivo,
// o si el cliente no existe.
func TestCP2_R1_ReferenciasValidas(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Venta{MedicamentoID: 99999, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un medicamento inexistente debe devolver ErrReferenciaInvalida")

	extra := models.Medicamento{Nombre: "Jarabe descontinuado", PrecioUnitario: 15, Stock: 3, Activo: false}
	require.NoError(t, e.medicamentos.Crear(&extra))
	a = models.Venta{MedicamentoID: extra.ID, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un medicamento INACTIVO debe devolver ErrReferenciaInvalida")

	a = models.Venta{MedicamentoID: e.principal.ID, ClienteID: 99999, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un cliente inexistente debe devolver ErrReferenciaInvalida")
}

// R2: la cantidad no puede superar el stock disponible.
func TestCP2_R2_StockInsuficiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Venta{MedicamentoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 11}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrStockInsuficiente,
		"pedir 11 unidades con stock 10 debe devolver ErrStockInsuficiente")
}

// R3: Total = Cantidad x PrecioUnitario, con 10% de descuento desde 5 unidades.
func TestCP2_R3_CalculoTotal(t *testing.T) {
	e := nuevoEntornoReglas(t)

	sinDescuento := models.Venta{MedicamentoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&sinDescuento),
		"crear una venta válida no debe devolver error")
	require.InDelta(t, 25.50, sinDescuento.Total, 0.001,
		"3 x 8.50 = 25.50 (sin descuento)")
	require.Equal(t, models.EstadoPendiente, sinDescuento.Estado,
		"una venta recién creada debe quedar en estado PENDIENTE")

	conDescuento := models.Venta{MedicamentoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 5}
	require.NoError(t, e.svc.Crear(&conDescuento))
	require.InDelta(t, 38.25, conDescuento.Total, 0.001,
		"5 x 8.50 = 42.50, con 10% de descuento = 38.25")
}

// R4: solo se puede anular una venta en estado PENDIENTE.
func TestCP2_R4_AnularSoloPendiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	despachada := models.Venta{
		MedicamentoID: e.principal.ID,
		ClienteID:     e.ana.ID,
		Cantidad:      1,
		Estado:        models.EstadoDespachada,
		Total:         8.5,
	}
	require.NoError(t, e.ventas.Crear(&despachada))
	require.ErrorIs(t, e.svc.Anular(despachada.ID), services.ErrEstadoInvalido,
		"anular una venta DESPACHADA debe devolver ErrEstadoInvalido")

	require.ErrorIs(t, e.svc.Anular(99999), services.ErrNoEncontrado,
		"anular una venta inexistente debe devolver ErrNoEncontrado")
}

// R5: al crear se descuenta el stock; al anular, se repone.
func TestCP2_R5_ReposicionStock(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Venta{MedicamentoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&a))

	h, ok := e.medicamentos.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(7), h.Stock,
		"al crear una venta de 3 unidades, el stock debe bajar de 10 a 7")

	require.NoError(t, e.svc.Anular(a.ID), "anular una venta PENDIENTE debe funcionar")

	anulada, ok := e.ventas.ObtenerPorID(a.ID)
	require.True(t, ok)
	require.Equal(t, models.EstadoAnulada, anulada.Estado,
		"tras anular, la venta debe quedar en estado ANULADA")

	h, ok = e.medicamentos.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(10), h.Stock,
		"al anular, las 3 unidades deben reponerse al stock (7 -> 10)")
}
