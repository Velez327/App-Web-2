package services

import (
	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/storage"
)

// TAREA (CP2): Implemente VentaService con las 5 reglas de negocio.
//
// Las reglas están A LA VISTA en las pantallas (carpeta pantallas/) y los
// tests de acceptance/reglas_test.go las verifican una por una. Devuelva los
// errores de dominio de errores.go: los tests los comprueban con errors.Is.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Observe que el service recibe TRES repositories: necesita consultar
//     Medicamento y Cliente para validar, y actualizar Medicamento al anular.
type VentaService struct {
	ventas   storage.VentaRepository
	medicamentos storage.MedicamentoRepository
	clientes     storage.ClienteRepository
}

func NuevaVentaService(
	ventas storage.VentaRepository,
	medicamentos storage.MedicamentoRepository,
	clientes storage.ClienteRepository,
) *VentaService {
	return &VentaService{
		ventas:   ventas,
		medicamentos: medicamentos,
		clientes:     clientes,
	}
}

// Crear registra un nuevo venta aplicando R1, R2 y R3.
// TODO (R1): el medicamento debe existir y estar activo; el cliente debe existir.
// TODO (R2): la cantidad no puede superar el stock disponible del medicamento.
// TODO (R3): calcule el total (observe en las pantallas cuándo aplica descuento).
// TODO: al crear, el stock del medicamento se descuenta (mire la pantalla 01
// antes y después de crear una venta; R5 es la operación inversa).
func (s *VentaService) Crear(a *models.Venta) error {
	// R1: verificar que el medicamento existe y está activo
	med, ok := s.medicamentos.ObtenerPorID(a.MedicamentoID)
	if !ok || !med.Activo {
		return ErrReferenciaInvalida
	}

	// R1: verificar que el cliente existe
	_, ok = s.clientes.ObtenerPorID(a.ClienteID)
	if !ok {
		return ErrReferenciaInvalida
	}

	// R2: verificar stock suficiente
	if a.Cantidad > med.Stock {
		return ErrStockInsuficiente
	}

	// R3: calcular total con descuento desde 5 unidades
	total := float64(a.Cantidad) * med.PrecioUnitario
	if a.Cantidad >= 5 {
		total = total * 0.90
	}
	a.Total = total
	a.Estado = models.EstadoPendiente

	// R5: descontar stock
	med.Stock -= a.Cantidad
	if err := s.medicamentos.Actualizar(&med); err != nil {
		return err
	}

	return s.ventas.Crear(a)
}

func (s *VentaService) ObtenerPorID(id uint) (models.Venta, error) {
	a, ok := s.ventas.ObtenerPorID(id)
	if !ok {
		return models.Venta{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *VentaService) Listar() ([]models.Venta, error) {
	// TODO: implementar.
	return s.ventas.Listar()
}

// Anular cancela una venta aplicando R4 y R5.
// TODO (R4): solo se puede anular una venta en estado PENDIENTE.
// TODO (R5): al anular, la cantidad se repone al stock del medicamento.
func (s *VentaService) Anular(id uint) error {
	// TODO: implementar.
	// R4: verificar que la venta existe
	a, ok := s.ventas.ObtenerPorID(id)
	if !ok {
		return ErrNoEncontrado
	}

	// R4: solo se puede anular si está PENDIENTE
	if a.Estado != models.EstadoPendiente {
		return ErrEstadoInvalido
	}

	// R5: reponer stock
	med, ok := s.medicamentos.ObtenerPorID(a.MedicamentoID)
	if !ok {
		return ErrReferenciaInvalida
	}
	med.Stock += a.Cantidad
	if err := s.medicamentos.Actualizar(&med); err != nil {
		return err
	}

	a.Estado = models.EstadoAnulada
	return s.ventas.Actualizar(&a)
}
