// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-farmacia/internal/models"

// VentaRepository define el contrato de persistencia de Venta.
// Su implementación GORM (en venta_gorm.go) debe satisfacer EXACTAMENTE
// estas firmas. Observe que el repositorio NO contiene lógica de negocio:
// las reglas (validaciones, cálculo del total, anulación) viven en el service.
type VentaRepository interface {
	Crear(a *models.Venta) error
	ObtenerPorID(id uint) (models.Venta, bool)
	Listar() ([]models.Venta, error)
	Actualizar(a *models.Venta) error
}
